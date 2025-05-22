package platform

import (
	"fmt"
	"spider/util"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Jingdong struct {
	filepath string
	wg       *sync.WaitGroup
}

func NewJingdong(filepath string, wg *sync.WaitGroup) *Jingdong {
	return &Jingdong{filepath: filepath, wg: wg}
}

func (jd *Jingdong) Start() {

	jd.wg.Add(1)
	go func() {
		defer jd.wg.Done()
		jd.Run("./conf/jingdong/slide.txt", "slide")
	}()

	jd.wg.Add(1)
	go func() {
		defer jd.wg.Done()
		jd.Run("./conf/jingdong/pic.txt", "pic")
	}()

	jd.wg.Add(1)
	go func() {
		defer jd.wg.Done()
		jd.Run("./conf/jingdong/detail.txt", "detail")
	}()
	jd.wg.Wait()
}

func (jd *Jingdong) Run(filename string, typeName string) {
	content, err := util.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		panic(fmt.Sprintf("解析HTML失败: %v", err))
	}

	var imgURLs []string
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		imgURL, exists := s.Attr("src")
		if exists && imgURL != "" {
			imgURLs = append(imgURLs, jd.imgUrlFormat(imgURL))
		}
	})

	util.Download(jd.filepath, typeName, imgURLs)
}

func (jd *Jingdong) imgUrlFormat(imgURL string) string {
	if !strings.Contains(imgURL, "https:") {
		imgURL = "https:" + imgURL
	}
	if strings.Contains(imgURL, ".avif") {
		imgURL = strings.Split(imgURL, ".avif")[0]
	}
	if strings.Contains(imgURL, "s114x114_jfs") {
		imgURL = strings.ReplaceAll(imgURL, "s114x114_jfs", "s800x800_jfs")
	}
	if strings.Contains(imgURL, "s28x28_jfs") {
		imgURL = strings.ReplaceAll(imgURL, "s28x28_jfs", "s800x800_jfs")
	}

	return imgURL
}

package platform

import (
	"fmt"
	"spider/util"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Taobao struct {
	filepath string
	wg       *sync.WaitGroup
}

func NewTaobao(filepath string, wg *sync.WaitGroup) *Taobao {
	return &Taobao{filepath: filepath, wg: wg}
}

func (tb *Taobao) Start() {

	tb.wg.Add(1)
	go func() {
		defer tb.wg.Done()
		tb.Run("./conf/taobao/slide.txt", "slide")
	}()

	tb.wg.Add(1)
	go func() {
		defer tb.wg.Done()
		tb.Run("./conf/taobao/pic.txt", "pic")
	}()

	tb.wg.Add(1)
	go func() {
		defer tb.wg.Done()
		tb.Run("./conf/taobao/detail.txt", "detail")
	}()
	tb.wg.Wait()
}

func (tb *Taobao) Run(filename string, typeName string) {
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
			imgURLs = append(imgURLs, util.ImgWebpReplace(imgURL))
		}
	})

	util.Download(tb.filepath, typeName, imgURLs)
}

package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(filename string) (content string, err error) {
	htmlContent, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer htmlContent.Close()

	contentBytes, err := io.ReadAll(htmlContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	content = string(contentBytes)
	return
}

func ImgWebpReplace(imgURL string) (res string) {
	res = imgURL
	if !strings.Contains(res, "https:") {
		res = "https:" + res
	}	
	if strings.Contains(res, ".jpg") {
		res = strings.Split(res, "jpg")[0] + "jpg"
	}
	if strings.Contains(res, ".png") {
		res = strings.Split(res, "png")[0] + "png"
	}
	if strings.Contains(res, ".jpeg") {
		res = strings.Split(res, "jpeg")[0] + "jpeg"
	} 
	
	return
}

func Download(path string, typeName string, imgURLs []string) {

	for idx, imgURL := range imgURLs {
		fileName := fmt.Sprintf("%s_%d.jpg", typeName, idx+1)

		// 下载图片
		fmt.Printf("正在下载: %s\n", imgURL)
		if err := downloadImage(imgURL, filepath.Join(path, fileName)); err != nil {
			fmt.Printf("下载失败 %s: %v\n", imgURL, err)
			continue
		}
		fmt.Printf("已保存: %s\n", fileName)
	}
}

func downloadImage(imgURL, savePath string) error {
	// 创建HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", imgURL, nil)
	if err != nil {
		return err
	}

	// 添加User-Agent头（防止被拒绝）
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode)
	}

	// 创建输出文件
	out, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 复制内容到文件
	_, err = io.Copy(out, resp.Body)
	return err
}

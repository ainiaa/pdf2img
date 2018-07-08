package main

import (
	"fmt"

	"maicai.ddxq.com/v3"
)

func main() {
	pdfName := "d:/www/150003521055_82538184.pdf"
	savePath := "d:/www/150003521055_dd.png"
	format := "png"
	width, height, err := pdf2img.ConvertToImg(pdfName, savePath, 180, 100, format)
	if err != nil {
		fmt.Errorf("pdf2img error:%s", err.Error())
	}
	fmt.Printf("width:%d height:%d ", width, height)
}

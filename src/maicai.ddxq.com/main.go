package main

import (
	"fmt"
	//"os"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	pdfName := "d:/www/150003521055_82538184.pdf"
	savePath := "d:/www/150003521055_82538184.png"
	format := "png"
	width, height, err := pdf2img(pdfName, savePath, 180, 100, format)
	if err != nil {
		fmt.Errorf("pdf2img error:%s", err.Error())
	}
	fmt.Printf("width:%d height:%d ", width, height)
}

func pdf2img(pdfName string, savePath string, resolution float64, compressionQuality uint, format string) (width uint, height uint, err error) {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	//mw.SetImageCompressionQuality(compressionQuality)
	mw.SetResolution(resolution, resolution)
	if err := mw.ReadImage(pdfName); err != nil {
		fmt.Printf("文件读取失败! error:%s", err.Error())
		fmt.Println()
	}

	pages := int(mw.GetNumberImages())

	fmt.Printf("GetNumberImages:%d", pages)
	fmt.Println()

	newMw := imagick.NewMagickWand()
	defer newMw.Destroy()

	for i := 0; i < pages; i++ {
		mw.SetIteratorIndex(i)
		//压平图像，去掉alpha通道，防止JPG中的alpha变黑,用在ReadImage之后
		if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_OFF); err != nil {
			fmt.Println("压平图像，去掉alpha通道，防止JPG中的alpha变黑")
			return 0, 0, err
		}

		tmpMw := imagick.NewMagickWandFromImage(mw.GetImageFromMagickWand())
		tmpMw.SetImageCompression(imagick.COMPRESSION_NO)
		tmpMw.SetImageCompressionQuality(compressionQuality)
		tmpMw.SetFormat(format)
		tmpMw.StripImage()
		tmpMw.TrimImage(0)
		tmpWidth := tmpMw.GetImageWidth() + 10
		tmpHeight := tmpMw.GetImageHeight() + 10
		if i+1 == pages {
			tmpHeight += 10
		}

		newMw.NewImage(tmpWidth, tmpHeight, imagick.NewPixelWand())
		newMw.CompositeImage(tmpMw, imagick.COMPOSITE_OP_PLUS, false, 5, 5)

	}
	newMw.ResetIterator() //不重置 就会变成一张图片
	if err = newMw.AppendImages(true).WriteImage(savePath); err != nil {
		fmt.Printf("save png err:%s", err.Error())
	}
	return newMw.GetImageWidth(), newMw.GetImageHeight(), err
}

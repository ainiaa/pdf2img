package pdf2img

import (
	"fmt"
	"os"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func printEnv() {
	environ := os.Environ()
	for i := range environ {
		fmt.Println(environ[i])
	}
}

func ConvertToImg(pdfName string, savePath string, resolution float64, compressionQuality uint, format string) (width uint, height uint, err error) {

	//printEnv()

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	if isExists, err := PathExists(pdfName); !isExists {
		if err != nil {
			fmt.Printf("文件%s不存在 error：%s", pdfName, err.Error())
		} else {
			fmt.Printf("文件%s不存在", pdfName)
		}

		fmt.Println()
	} else {
		fmt.Printf("文件%s存在", pdfName)
		fmt.Println()
	}

	//mw.SetImageCompressionQuality(compressionQuality)
	mw.SetResolution(resolution, resolution)
	if err := mw.ReadImage(pdfName); err != nil {
		fmt.Printf("文件读取失败! error:%s", err.Error())
		fmt.Println()
	} else {
		fmt.Printf("文件读取成功")
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

		newMw.NewImage(tmpMw.GetImageWidth(), tmpMw.GetImageHeight(), imagick.NewPixelWand())
		newMw.CompositeImage(tmpMw, imagick.COMPOSITE_OP_SRC, true, 0, 0)
		//newMw.CompositeImage(tmpMw, imagick.COMPOSITE_OP_PLUS, false, 5, 5)

	}
	newMw.ResetIterator() //不重置 就会变成一张图片
	if err = newMw.AppendImages(true).WriteImage(savePath); err != nil {
		fmt.Printf("save png err:%s", err.Error())
	}

	fmt.Println("convert finish!!")
	return newMw.GetImageWidth(), newMw.GetImageHeight(), err
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

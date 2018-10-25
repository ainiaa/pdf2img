package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"maicai.ddxq.com/v3"
)

//ConverRequester 请求数据结构
// PdfName: pdf文件名称
// Format：需要format的格式 png jpeg等
// SavePath：保存路径
// CompressionQuality：压缩质量
type ConverRequester struct {
	PdfName            string  `json:"pdfname,omitemty"`            //pdf文件名称
	Format             string  `json:"format,omitemty"`             //图片格式
	SavePath           string  `json:"savepath,omitemty"`           //图片保存路径
	Resolution         float64 `json:"resolution,omitemty"`         //图片质量
	CompressionQuality uint    `json:"compressionQuality,omitemty"` //图片压缩比
}

// ConvertResponser 相应内容
// Status：状态码
// Width：图片宽度
// Height：图片高度
// Size：图片大小
// Message：消息内容
type ConvertResponser struct {
	Status  int    `json:"status,omitemty"`
	Width   uint   `json:"width,omitemty"`
	Height  uint   `json:"height,omitemty"`
	Size    int    `json:"size,omitemty"`
	Message string `json:"message,omitemty"`
}

// MIN 最小值
const MIN = 0.000001

// IsEqual 是否相等
func IsEqual(f1, f2 float64) bool {
	return math.Dim(f1, f2) < MIN
}

// PdfConvertToImg pdf转为png
func PdfConvertToImg(w http.ResponseWriter, req *http.Request) {

	req.ParseForm() //解析form内容

	pdfName := req.PostForm.Get("pdfname")
	savePath := req.PostForm.Get("savepath")
	resolution, _ := strconv.ParseFloat(req.PostForm.Get("resolution"), 64)
	compressionQuality, _ := strconv.Atoi(req.PostForm.Get("compressionquality"))
	format := req.PostForm.Get("format")
	if IsEqual(resolution, 0.0) {
		resolution = 180.0
	}
	if compressionQuality == 0 {
		compressionQuality = 100
	}
	convertRes := ConvertResponser{}
	fmt.Printf("pdfName:%s", pdfName)
	fmt.Println("before call ")
	if pdfName != "" {
		width, height, err := pdf2img.ConvertToImg(pdfName, savePath, resolution, uint(compressionQuality), format)
		fmt.Printf("width:%d height:%d ", int(width), int(height))
		if err != nil {
			convertRes.Status = 0
			convertRes.Message = err.Error()
		} else {
			convertRes.Status = 1
			convertRes.Message = "success"
			convertRes.Width = width
			convertRes.Height = height
		}
		fmt.Printf("width:%d height:%d", int(width), int(height))
	}

	json.NewEncoder(w).Encode(convertRes)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/pdf2img", PdfConvertToImg).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}

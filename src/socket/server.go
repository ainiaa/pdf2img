package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	"os"

	"maicai.ddxq.com/v3"
)

//请求数据结构
type ConverRequester struct {
	PdfName            string  `json:pdfname,omitemty`
	Format             string  `json:format,omitemty`
	SavePath           string  `json:savepath,omitemty`
	Resolution         float64 `json:resolution,omitemty`
	CompressionQuality uint    `json:compressionQuality,omitemty`
}

const MIN = 0.000001

func IsEqual(f1, f2 float64) bool {
	return math.Dim(f1, f2) < MIN
}

func main() {

	//建立socket，监听端口
	netListen, err := net.Listen("tcp", "127.0.0.1:9999")
	CheckError(err)
	defer func(l net.Listener) {
		fmt.Println("关闭")
		l.Close()
	}(netListen)
	Log("Waiting for clients")
	for {

		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " 连接成功请求地址")
		go handleConnection(conn)
	}

}

//处理连接
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	Log("走了处理请求")
	for {
		Log("走的次数")
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), "连接错误的请求地址: ", err)
			return
		}
		Log(conn.RemoteAddr().String(), "这是啥数据:\n", string(buffer[:n]))
		if len(string(buffer[:n])) > 25 {
			sender(conn, string(buffer[:n])) //直接使用buffer 会报错 json.Unmarshal error: invalid character '\x00' after top-level value
		}
	}
}
func sender(conn net.Conn, content string) {
	Log("需要发送的xml")
	var buffer bytes.Buffer
	/*
		//var sl []string
		buffer.WriteString("<?xml version=\"1.0\" encoding=\"GBK\"?>")
		buffer.WriteString("<message>")
		buffer.WriteString("<head>")
		buffer.WriteString("<field name=\"ReceiveTime\">112823</field>")
		buffer.WriteString("<field name=\"ReceiveDate\">20151101</field>")
		buffer.WriteString("</head>")
		buffer.WriteString("<body>")
		buffer.WriteString("<field name=\"Host\">20151101</field>")
		buffer.WriteString("</body>")
		buffer.WriteString("</message>")
		Log(buffer.Bytes())
		Log("地址为===" + conn.RemoteAddr().String())
		//conn.Write([]byte(strings.Join(sl, "")))
		//-->使用数组的形式 得到byte也行 只不过看着没buffer这样的好
		// ar := []byte {1, 1,1, 1}
		//for i:= 0;i< len(buffer.Bytes()); i++ {
		//    ar = append(ar,buffer.Bytes()[i])
		//}
		//Log(ar)

		Log(buffer.String())
		conn.Write(buffer.Bytes())
		Log("send over")
	*/

	requester := &ConverRequester{}
	err := json.Unmarshal([]byte(content), requester)
	if err != nil {
		Log("json.Unmarshal error:", err.Error())
	}

	pdfNameLog := fmt.Sprintf("pdfName:%s", requester.PdfName)
	Log(pdfNameLog)
	resolution := requester.Resolution
	compressionQuality := requester.CompressionQuality
	format := requester.Format
	if !IsEqual(resolution, 0.0) {
		resolution = 180.0
	}
	if compressionQuality != 0 {
		compressionQuality = 100
	}
	width, height, err := pdf2img.ConvertToImg(requester.PdfName, requester.SavePath, resolution, compressionQuality, format)
	if err != nil {
		fmt.Printf("ConvertToImg failure error:%s", err.Error())
		buffer.WriteString("ConvertToImg failure error:")
		buffer.WriteString(err.Error())
	} else {
		s := fmt.Sprintf("ConvertToImg success:width:%d, height:%d", width, height)
		fmt.Println(s)
		buffer.WriteString("ConvertToImg success")
		buffer.WriteString(s)
	}

	conn.Write(buffer.Bytes())
	Log("send over")

}
func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

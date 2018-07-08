package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//请求数据结构
type ConverRequester struct {
	PdfName            string  `json:pdfname,omitemty`
	Format             string  `json:format,omitemty`
	SavePath           string  `json:savepath,omitemty`
	Resolution         float64 `json:resolution,omitemty`
	CompressionQuality uint    `json:compressionQuality,omitemty`
}

func sender(conn net.Conn) {
	requester := &ConverRequester{}
	requester.PdfName = "E:/tmp/150003521055_82538184.pdf"
	requester.SavePath = "E:/tmp/150003521055_ddd.png"
	data, _ := json.Marshal(requester)

	conn.Write(data)
	fmt.Println("send over")
}

func main() {
	server := "127.0.0.1:9999"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	sender(conn)
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	fmt.Println(string(buffer[:n]))
}

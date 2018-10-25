package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main2() {

	url := "http://localhost:12345/people/1"

	payload := strings.NewReader("{\n  \"firstname\": \"wang\",\n  \"lastname\": \"shubo\",\n  \"address\": {\n    \"city\": \"Beijing\",\n    \"state\": \"Beijng\"\n  }\n}")

	req, _ := http.NewRequest("DELETE", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "4a894ad6-2887-259a-c953-5d26fed70963")

	res, _ := http.DefaultClient.Do(req)

	if res != nil {
		defer res.Body.Close()
	}

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

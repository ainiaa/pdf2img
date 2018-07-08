package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:12345/people/3"

	payload := strings.NewReader("{\n  \"firstname\": \"wang\",\n  \"lastname\": \"shubo\",\n  \"address\": {\n    \"city\": \"Beijing\",\n    \"state\": \"Beijng\"\n  }\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "a9d590dd-1819-15f6-962e-0eabf4b7e707")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

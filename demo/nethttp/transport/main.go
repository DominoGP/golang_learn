package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 1. 创建 Transport
	transport := &http.Transport{}

	// 2. 使用 Transport 发送请求
	req, err := http.NewRequest("GET", "https://www.baidu.com", nil)
	resp, err := transport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 3. 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 响应内容
	fmt.Println(string(body))
}

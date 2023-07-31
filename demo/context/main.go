package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 创建一个带有超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 使用context创建一个http请求
	req, err := http.NewRequestWithContext(ctx, "GET", "http://baidu.com", nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 创建一个http客户端并发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 处理响应
	fmt.Println("响应状态码:", resp.StatusCode)
}

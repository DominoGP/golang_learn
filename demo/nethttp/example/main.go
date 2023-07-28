package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 一个处理请求的handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	// 创建一个默认的路由
	mux := http.NewServeMux()
	// 为路径挂载对应的处理器
	mux.HandleFunc("/hello", helloHandler)

	// 创建服务器，指定对应的端口和处理路由
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// 启动服务
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

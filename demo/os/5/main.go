package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// signalTest()
	printTime()
}

func signalTest() {
	// 监听中断信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("程序已启动")

	// 监听信号chan
	for {
		select {
		case <-sigChan:
			fmt.Println("接收到中断信号，程序即将退出")
			return
		default:
			fmt.Println("程序正在运行")
			time.Sleep(time.Second)
		}
	}
}

func printTime() {
	// 监听中断信号和SIGTERM信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// 等待信号
	fmt.Println("等待信号...")
	sig := <-sigChan
	fmt.Println("接收到信号：", sig)

	// 输出当前时间
	fmt.Println("当前时间：", time.Now())

	// 程序即将退出
	fmt.Println("程序即将退出")
}

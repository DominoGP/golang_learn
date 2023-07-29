package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// 设置环境变量
	os.Setenv("NAME", "John")

	// 获取环境变量
	value := os.Getenv("NAME")
	fmt.Println(value) // John

	// 移除环境变量
	os.Unsetenv("NAME")

	// 获取所有环境变量
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], pair[1])
	}

	// 清除所有环境变量
	os.Clearenv()
}

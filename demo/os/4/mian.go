package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
)

func main() {

	// 获取主机名
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(host)

	// 获取用户信息
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("User: %s\n", user.Name)

	// 获取CPU数
	numCPU := runtime.NumCPU()
	fmt.Println("CPU: ", numCPU)

	// 获取GOMAXPROCS
	maxprocs := runtime.GOMAXPROCS(0)
	fmt.Println("GOMAXPROCS: ", maxprocs)

	// 获取当前程序启动参数
	arg := os.Args
	fmt.Println(arg)
}

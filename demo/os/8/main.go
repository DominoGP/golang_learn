package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func main() {

	// 返回系统的CPU数
	cpuNum := os.Getenv("GOMAXPROCS")
	if cpuNum == "" {
		cpuNum = strconv.Itoa(runtime.NumCPU())
	}
	fmt.Printf("CPU number: %s\n", cpuNum)

	// 设置最大使用的CPU数
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 返回cgo调用次数
	fmt.Printf("Number of cgo calls: %d\n", runtime.NumCgoCall())

}

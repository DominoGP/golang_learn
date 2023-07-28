package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("io/files/data.txt")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	_, err = file.Seek(10, 0) // 从文件起始位置向后移动10个字节
	if err != nil {
		fmt.Println("Failed to seek:", err)
		return
	}

	buffer := make([]byte, 5) // 创建一个5个字节的切片
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Failed to read:", err)
		return
	}

	fmt.Println(string(buffer))
}

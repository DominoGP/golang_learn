package main

import (
	"fmt"
	"os"
)

func main() {

	// 创建文件
	file, err := os.Create("./filename.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//获取文件名称

	fileInfo, err := os.Stat("./filename.txt")
	fmt.Println(fileInfo.Name())    //获取文件名
	fmt.Println(fileInfo.IsDir())   //判断是否是目录，返回bool类型
	fmt.Println(fileInfo.ModTime()) //获取文件修改时间
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.Size()) //获取文件大小
	fmt.Println(fileInfo.Sys())

	fmt.Println()

	// 写入文件内容
	file.WriteString("contents")

	// 读取文件内容
	data, err := os.ReadFile("./filename.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	//获取当前的工作目录的根路径
	fmt.Println(os.Getwd())

	// 重命名文件
	err = os.Rename("./filename.txt", "./new.txt")
	if err != nil {
		fmt.Println(err) // 打印详细错误信息
	}
	fmt.Println("ok")

	// 修改文件权限
	err = os.Chmod("./new.txt", 0666)

	// 创建目录
	err = os.Mkdir("./subdir", 0755)

	// 读取目录
	files, err := os.ReadDir("./subdir")
	fmt.Println(files)

	// 删除文件
	err = os.Remove("./filename.txt")

	// 删除目录
	err = os.Remove("./subdir")
}

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	stdioTest()
	cmdTest()
	channelTest()
}

func stdioTest() {
	// 从标准输入读取数据
	fmt.Print("请输入内容：")
	var input string
	fmt.Scanln(&input)

	// 将数据写入标准输出和标准错误输出
	fmt.Fprintln(os.Stdout, "这是标准输出：", input)
	fmt.Fprintln(os.Stderr, "这是标准错误输出：", input)
}

func cmdTest() {
	// 创建一个命令
	cmd := exec.Command("ls", "-l")

	// 打开一个文件，用于保存命令的输出
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("创建文件失败：", err)
		return
	}

	// 将命令的标准输出和标准错误输出重定向到文件
	cmd.Stdout = file
	cmd.Stderr = file

	// 启动命令
	err = cmd.Start()
	if err != nil {
		fmt.Println("启动命令失败：", err)
		return
	}

	// 等待命令执行完毕
	err = cmd.Wait()
	if err != nil {
		fmt.Println("等待命令执行失败：", err)
		return
	}
	fmt.Println("命令执行完成")
}

func channelTest() {
	// 创建一个管道
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Println("创建管道失败：", err)
		return
	}

	// 启动一个协程，向管道中写入数据
	go func() {
		defer writer.Close()
		fmt.Fprintln(writer, "这是一条来自子进程的消息")
	}()

	// 从管道中读取数据
	buf := make([]byte, 1024)
	n, err := reader.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("读取数据失败：", err)
		return
	}
	fmt.Println("收到一条来自子进程的消息：", string(buf[:n]))
}

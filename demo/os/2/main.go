package main

import (
	"fmt"
	"os"
)

func main() {

	// 获取进程ID
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)

	// 获取父进程ID
	ppid := os.Getppid()
	fmt.Println("Parent Process ID:", ppid)

	// 退出进程
	os.Exit(0)

	// 杀死进程(linux)
	// os.Kill(pid, syscall.SIGKILL)
	// windows
	// os.Kill(pid, os.Interrupt)

	// 获取用户ID
	uid := os.Getuid()
	fmt.Println("User ID:", uid)

	// 设置用户ID需要root权限 linux
	// os.Setuid(1000)
}

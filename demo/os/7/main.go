package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	// 获取当前用户信息
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("获取当前用户信息失败：", err)
		return
	}

	fmt.Println("当前用户信息：", currentUser.Username, currentUser.HomeDir)

	// 根据用户名获取用户信息
	userInfo, err := user.Lookup("root")
	if err != nil {
		fmt.Println("获取用户信息失败：", err)
		return
	}

	fmt.Println("用户信息：", userInfo.Username, userInfo.Uid, userInfo.HomeDir)

	// 根据用户ID获取用户信息
	userInfo, err = user.LookupId("0")
	if err != nil {
		fmt.Println("获取用户信息失败：", err)
		return
	}

	fmt.Println("用户信息：", userInfo.Username, userInfo.Uid, userInfo.HomeDir)

	// 根据群组名获取群组信息
	groupInfo, err := user.LookupGroup("root")
	if err != nil {
		fmt.Println("获取群组信息失败：", err)
		return
	}

	fmt.Println("群组信息：", groupInfo.Name, groupInfo.Gid)

	// 根据群组ID获取群组信息
	groupInfo, err = user.LookupGroupId("0")
	if err != nil {
		fmt.Println("获取群组信息失败：", err)
		return
	}

	fmt.Println("群组信息：", groupInfo.Name, groupInfo.Gid)

	// 打开文件
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("创建文件失败：", err)
		return
	}
	defer file.Close()

	// 设置文件所有者
	err = file.Chown(0, 0)
	if err != nil {
		fmt.Println("设置文件所有者失败：", err)
		return
	}

	// 设置文件权限
	err = file.Chmod(0644)
	if err != nil {
		fmt.Println("设置文件权限失败：", err)
		return
	}

	fmt.Println("文件所有者和权限设置完成")
}

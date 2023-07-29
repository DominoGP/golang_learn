package main

import (
	"fmt"
	"time"
)

func main() {

	// 1. 获取当前时间
	now := time.Now()

	// 2. 暂停程序
	time.Sleep(2 * time.Second)

	// 3. 格式化时间
	format := now.Format("2006/01/02 15:04:05")
	fmt.Println(format)

	// 4. 时区转换
	loc, _ := time.LoadLocation("Local")
	locTime := now.In(loc)
	fmt.Println(locTime)

	// 5. 解析时间戳
	timestamp := now.Unix()
	t := time.Unix(timestamp, 0)

	// 6. 日期计算
	t = t.Add(24 * time.Hour)

	// 7. 比较时间
	t1 := time.Now()
	if t1.Before(t) {

	}

	// 8. 定时器
	timer := time.NewTimer(5 * time.Second)
	<-timer.C // 等待5秒

	// 9. 经过时间
	duration := time.Since(t1)
	fmt.Println(duration)

	// 10. 本地时间
	fmt.Println(time.Now().Local())

	// 11. 常量
	second := time.Second
	fmt.Println(second)
}

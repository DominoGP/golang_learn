package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// basicLog()

	//设置日志输出层级为LevelInfo，小于该层级的日志皆不输出
	SetLogLevel(LevelInfo)

	Log(LevelDebug, "debug")
	Log(LevelInfo, "info")
	Log(LevelWarn, "warn")

	//设置日志输出层级为LevelWarn，小于该层级的日志皆不输出
	SetLogLevel(LevelWarn)

	Log(LevelDebug, "debug")
	Log(LevelInfo, "info")
	Log(LevelWarn, "warn")
}

func basicLog() {
	log.Print("This is a log message")
	log.Printf("This is a formatted log message: %s", "Hello, World!")
	log.Println("This is a log message with a newline")

	log.Fatal("This is a fatal error log message")
	log.Fatalf("This is a formatted fatal error log message: %s", "Hello, World!")
	log.Fatalln("This is a fatal error log message with a newline")

	log.Panic("This is a panic log message")
	log.Panicf("This is a formatted panic log message: %s", "Hello, World!")
	log.Panicln("This is a panic log message with a newline")
}

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// 初始化globalLogLevel参数，避免空值
var globalLogLevel = LevelInfo

// 封装日志输出函数
func Log(level int, v ...any) {
	if level >= globalLogLevel {
		// log.Println(v...)

		f, _ := os.OpenFile("log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		fmt.Fprintln(f, v)
	}
}

// 设置日志输出级别函数
func SetLogLevel(level int) {
	globalLogLevel = level
}

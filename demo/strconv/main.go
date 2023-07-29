package main

import (
	"fmt"
	"strconv"
)

func main() {
	basicType()
	otherFunc()
}

func basicType() {
	a, _ := strconv.Atoi("123456")
	fmt.Printf("a values %v\n", a)
	b, _ := strconv.ParseInt("123456", 10, 32)
	fmt.Printf("b values %v\n", b)
	c, _ := strconv.ParseInt("123456", 10, 16)
	fmt.Printf("c values %v\n", c)
	d, _ := strconv.ParseInt("111", 2, 32)
	fmt.Printf("d values %v\n", d)

	var i int = 12345
	s := strconv.Itoa(i)
	fmt.Printf("s values %v,type is %T", s, s)
}
func otherFunc() {
	// 将整数类型追加到字节数组中
	b1 := make([]byte, 0, 100)
	b1 = strconv.AppendInt(b1, 123, 10)
	fmt.Printf("Type: %T, Value: %v\n", b1, string(b1))

	// 将无符号整数类型追加到字节数组中
	b2 := make([]byte, 0, 100)
	b2 = strconv.AppendUint(b2, 456, 10)
	fmt.Printf("Type: %T, Value: %v\n", b2, string(b2))

	// 将布尔值类型追加到字节数组中
	b3 := make([]byte, 0, 100)
	b3 = strconv.AppendBool(b3, true)
	fmt.Printf("Type: %T, Value: %v\n", b3, string(b3))

	// 将浮点数类型追加到字节数组中
	b4 := make([]byte, 0, 100)
	b4 = strconv.AppendFloat(b4, 3.14, 'f', 2, 64)
	fmt.Printf("Type: %T, Value: %v\n", b4, string(b4))

	// 将字符串转换为带引号的字符串
	s1 := strconv.Quote("hello world")
	fmt.Printf("Type: %T, Value: %v\n", s1, s1)

	// 将带引号的字符串转换为普通字符串
	s2, err := strconv.Unquote(`"hello world"`)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("Type: %T, Value: %v\n", s2, s2)

	// 将 rune 类型的字符转换为带引号的字符串
	s3 := strconv.QuoteRune('中')
	fmt.Printf("Type: %T, Value: %v\n", s3, s3)

	// 将带引号的字符串转换为 rune 类型的字符
	r, _, _, err := strconv.UnquoteChar(`'\u4e2d'`, '\'')
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("Type: %T, Value: %v\n", r, r)
}

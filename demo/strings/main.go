package main

import (
	"fmt"
	"strings"
)

func main() {
	basicFunc()
	otherFunc()
}
func basicFunc() {
	// 判断字符串是否包含另一个字符串
	fmt.Println(strings.Contains("hello world", "world"))

	// 查找字符串中某个子串的位置
	fmt.Println(strings.Index("hello world", "world"))

	// 将字符串中的某个子串替换为另一个字符串
	fmt.Println(strings.Replace("hello world", "world", "golang", -1))

	// 将字符串按照指定的分隔符分割为多个子串
	fmt.Println(strings.Split("hello world", " "))

	// 将多个字符串连接起来，中间用指定的分隔符隔开
	fmt.Println(strings.Join([]string{"hello", "world"}, " "))

	// 去掉字符串开头和结尾的指定字符
	fmt.Println(strings.Trim(" hello ", " "))

	// 将字符串转换为大写形式
	fmt.Println(strings.ToUpper("hello world"))

	// 比较两个字符串的大小关系
	fmt.Println(strings.Compare("hello world", "Hello World"))
}

func otherFunc() {
	// 根据给定的函数将字符串分割成多个子串
	f := func(c rune) bool {
		return c == ' ' || c == ','
	}
	fmt.Println(strings.FieldsFunc("hello, world", f))

	// 根据给定的函数查找字符串中某个字符的位置
	f2 := func(c rune) bool {
		return c == 'o'
	}
	fmt.Println(strings.IndexFunc("hello world", f2))
	fmt.Println(strings.LastIndexFunc("hello world", f2))

	// 将字符串中的每个字符都映射到另一个字符
	f3 := func(c rune) rune {
		if c >= 'a' && c <= 'z' {
			return c - 'a' + 'A'
		}
		return c
	}
	fmt.Println(strings.Map(f3, "hello world"))

	// 将字符串包装在一个 io.Reader 接口中
	r := strings.NewReader("hello world")
	buf := make([]byte, 5)
	r.Read(buf)
	fmt.Println(string(buf))

	// 创建一个 strings.Replacer 对象，用于执行多个字符串替换操作
	r2 := strings.NewReplacer("hello", "hi", "world", "golang")
	fmt.Println(r2.Replace("hello world"))

	// 按照指定的后缀分隔字符串
	fmt.Println(strings.SplitAfter("hello,world,", ","))

	// 将字符串中的每个单词的首字母大写，并将所有字符转换为标题大小写
	fmt.Println(strings.Title("hello world"))
	fmt.Println(strings.ToTitle("hello world"))

	// 将字符串开头或结尾的所有满足指定条件的字符都删除
	f4 := func(c rune) bool {
		return c == ' '
	}
	fmt.Println(strings.TrimLeftFunc("   hello world   ", f4))
	fmt.Println(strings.TrimRightFunc("   hello world   ", f4))

	// 检查字符串是否包含有效的 UTF-8 字符,将无效的字符进行转换
	fmt.Println(strings.ToValidUTF8("hello world", " "))

}

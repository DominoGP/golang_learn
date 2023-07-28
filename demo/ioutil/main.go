package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	readAllTest()
	readerTest()
	readFileTest()

	writeFile()
	writeFile2()

	folerInfo()
}
func readAllTest() {
	file, err := os.Open("ioutil/files/test.txt")
	if err != nil {
		fmt.Println("open file err")
		return
	}

	c, err := ioutil.ReadAll(file)
	fmt.Println(err)
	fmt.Println(string(c))
}
func readerTest() {
	// NewReader创建一个从s读取数据的Reader。本函数类似bytes.NewBufferString
	reader := strings.NewReader("Hello word !")
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", all)
}

func readFileTest() {
	// 文件路径
	fileName := "ioutil/files/test.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", data)
}

func writeFile() {
	fileName := "ioutil/files/write.txt"
	for i := 0; i < 3; i++ {
		_ = ioutil.WriteFile(fileName, []byte("Hello World!"), os.ModePerm)
	}
	fmt.Println("write successful.")
}

func writeFile2() {
	var buffer bytes.Buffer

	for i := 0; i < 100; i++ {
		buffer.WriteString("this is line " + fmt.Sprintf("%d", i) + "\n")
	}
	if err := ioutil.WriteFile("ioutil/files/write2.txt", buffer.Bytes(), 0644); err != nil {
		fmt.Println(err)
	}
}

func folerInfo() {
	fileList, err := ioutil.ReadDir("ioutil")
	fmt.Println("\nfolder information:")
	for _, f := range fileList {
		fmt.Println(f.Name(), f.IsDir())
	}
	fmt.Println(err)
}

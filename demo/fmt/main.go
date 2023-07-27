package main

import "fmt"

func main() {
	// sout()
	stringOut()
}

func sout() {
	// 打印字符串
	fmt.Println("Hello, world!")

	// 打印整数
	num := 42
	fmt.Printf("The answer is %d\n", num)

	// 打印浮点数
	pi := 3.14159
	fmt.Printf("The value of pi is %.2f\n", pi)

	// 格式化输出
	name := "Alice"
	age := 25
	fmt.Printf("My name is %s and I'm %d years old.\n", name, age)

	// 格式化字符串
	message := fmt.Sprintf("The value of pi is %.2f", pi)
	fmt.Println(message)

	// 格式化输出
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// 格式化字符串
	str := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(str)

	// 格式化输入
	var inputName string
	var inputAge int
	fmt.Print("Enter your name: ")
	fmt.Scanln(&inputName)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&inputAge)
	fmt.Printf("Hello, %s! You are %d years old.\n", inputName, inputAge)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func stringOut() {
	person := Person{Name: "Alice", Age: 25}
	fmt.Println(person.String()) // 输出：Name: Alice, Age: 25
}

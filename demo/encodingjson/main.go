package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 编码为JSON
	person := Person{Name: "Alice", Age: 25}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return
	}
	fmt.Println(string(jsonData))

	// 解码JSON
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Println("JSON decoding error:", err)
		return
	}
	fmt.Println(decodedPerson)
}
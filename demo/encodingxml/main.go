package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	marshalDemo()
	unmarshalDemo()
}

func marshalDemo() {
	person := Person{
		Name: "Alice",
		Age:  30,
	}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		fmt.Println("XML encoding error:", err)
		return
	}

	fmt.Println(string(xmlData))
}
func unmarshalDemo() {
	xmlData := []byte(`
	<Person>
		<name>Alice</name>
		<age>30</age>
	</Person>
	`)

	var person Person
	err := xml.Unmarshal(xmlData, &person)
	if err != nil {
		fmt.Println("XML decoding error:", err)
		return
	}

	fmt.Println(person.Name, person.Age)
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	// getArgWithName()
	getArgWithoutName()
}

func getArgWithName() {
	flag.String("name", "Guest", "your name")
	flag.Int("age", 0, "your age")

	flag.Parse()

	name := flag.Arg(0)
	age := flag.Arg(1)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}

func getArgWithoutName() {
	flag.String("name", "Guest", "your name")
	flag.Int("age", 0, "your age")

	flag.Parse()

	name := flag.Arg(0)
	age := flag.Arg(1)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}

package main

import (
	"fmt"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("division by zero: %d / %d", x, y)
	}
	return x / y, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

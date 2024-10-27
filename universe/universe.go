package main

import (
	"fmt"
)

func checkNumber(num int) string {
	if num == 42 {
		return "Hello Universe"
	}
	return fmt.Sprintf("%d", num)
}

func main() {
	var input int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&input)

	result := checkNumber(input)
	fmt.Println(result)
}

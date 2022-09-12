package main

import (
	"fmt"
)

func BasicAtoi(s string) int {
	transform := []rune(s)
	for _, value := range transform {
		return value
	}
}
func main() {
	fmt.Println(BasicAtoi("12345"))
}

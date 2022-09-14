package piscine

import "fmt"

func Fibonacci(index int) int {
	if index == 0 {
		return 0
	}
	if index < 0 {
		return -1
	}
	if index == 1 {
		return 1
	}
	return index*Fibonacci(index-1) + Fibonacci(index-2)
}

func main() {
	arg1 := 1
	fmt.Println(Fibonacci(arg1))
}

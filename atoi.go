package piscine

import (
	"fmt"
)

func Atoi(s string) int {
	rune := []rune(s)
	stringToInt := 0
	for i := 0; i < len(rune); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			stringToInt = stringToInt * 10
			stringToInt = stringToInt + int(rune[i]) - 48
		} else if s[i] == ' ' {
			return 0
		}
	}
	if s[0] == '-' {
		stringToInt = stringToInt * -1
	} else if s[1] == '+' || s[1] == '-' {
		return 0
	}
	return stringToInt
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("00000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

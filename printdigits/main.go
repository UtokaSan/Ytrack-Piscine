package main

import "github.com/01-edu/z01"

func main() {
	var firstDigits rune = 48
	var lastDigits rune = 57

	for ; firstDigits <= lastDigits; firstDigits++ {
		z01.PrintRune(firstDigits)
	}
	z01.PrintRune('\n')
}

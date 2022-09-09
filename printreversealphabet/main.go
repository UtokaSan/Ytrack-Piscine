package main

import "github.com/01-edu/z01"

func main() {
	var firstLetter rune = 122
	var lastLetter rune = 97

	for ; firstLetter >= lastLetter; firstLetter-- {
		z01.PrintRune(firstLetter)
	}

	z01.PrintRune('\n')
}

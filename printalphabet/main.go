package main

import "github.com/01-edu/z01"

func main() {
	var firstLetter rune = 97
	var lastLetter rune = 122

	for ; firstLetter <= lastLetter; firstLetter++ {
		z01.PrintRune(firstLetter)
	}
	z01.PrintRune('\n')
}

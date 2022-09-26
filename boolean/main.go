package main

import (
	"github.com/01-edu/z01"
	"os"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
func even(nbr int) int {
	return nbr % 2
}
func isEven(nbr int) bool {
	if even(nbr) == 0 {
		return true
	}
	return false
}
func main() {
	lengthOfArg := len(os.Args[1])
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) == false {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := []rune(os.Args[1])

	for i := 0; i < len(args); i++ {
		z01.PrintRune(args[i])
	}
}

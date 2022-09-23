package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var nbNegOrPositif int = 49
	if n < 0 {
		nbNegOrPositif = 4549
		z01.PrintRune('-')
		n = n % 10 * nbNegOrPositif
		z01.PrintRune(rune(n))
	}
	if n > 0 {
		n = n % 10 * 1
		z01.PrintRune(rune(n))
	}
}

func main() {
	PrintNbr(1234)
	PrintNbr(-1234)
}

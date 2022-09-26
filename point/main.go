package main

import "github.com/01-edu/z01"

type point struct {
	x *int
	y *int
}

func setPoint(x *int, y *int) {
	*x = 42
	*y = 21
}

func main() {
	test := point{}
	x := 0
	y := 0
	setPoint(&x, &y)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(test.x)

}

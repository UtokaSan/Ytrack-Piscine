package piscine

import "fmt"

func TrimAtoi(s string) int {
	rune := []rune(s)
	a := 0
	b := 0
	c := 1
	d := false
	for _, word := range rune {
		if word >= '0' && word <= '9' {
			for i := '0'; i < word; i++ {
				a++
			}
			b = b*10 + a
			a = 0
		}
		if word == '-' && b == 0 {
			d = true
		}
	}
	if d == true {
		c = -1
	}
	return b * c
}

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}

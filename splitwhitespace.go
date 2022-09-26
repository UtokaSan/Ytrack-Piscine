package piscine

import "fmt"

func SplitWhiteSpaces(s string) []string {
	tab := []int
	for i := 0; i < len(s); i++ {
		tab = append(tab, i)
	}
	return tab
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
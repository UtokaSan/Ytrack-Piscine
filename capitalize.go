package piscine

import "fmt"

func Capitalize(s string) string {
	array := []rune(s)
	if array[0] >= 'a' && array[0] < 'z' {
		array[0] = rune(array[0] - 32)
	}
	return string(array)
}
func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

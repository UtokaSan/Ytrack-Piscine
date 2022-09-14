package piscine

func BasicAtoi(s string) int {
	rune := []rune(s)
	stringToInt := 0
	for i := 0; i < len(rune); i++ {
		stringToInt = stringToInt * 10
		stringToInt = stringToInt + int(rune[i]) - 48
	}
	return stringToInt
}

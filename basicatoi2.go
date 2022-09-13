package main

func BasicAtoi2(s string) int {
	rune := []rune(s)
	stringToInt := 0
	for i := 0; i < len(rune); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			stringToInt = stringToInt * 10
			stringToInt = stringToInt + int(rune[i]) - 48
		}
	}
	return stringToInt
}

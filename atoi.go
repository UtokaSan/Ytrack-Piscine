package piscine

func Atoi(s string) int {
	rune := []rune(s)
	stringToInt := 0
	for i := 0; i < len(rune); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			stringToInt = stringToInt * 10
			stringToInt = stringToInt + int(rune[i]) - 48
		}
		if s[i] == ' ' {
			return 0
		}
	}
	if s[0] == '-' {
		stringToInt = stringToInt * -1
	}
	if s[0] == '+' && s[1] == '+' || s[0] == '-' && s[1] == '-' {
		return 0
	}
	return stringToInt
}

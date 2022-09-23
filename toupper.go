package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for letter := range runes {
		if runes[letter] >= 'a' && runes[letter] <= 'z' {
			runes[letter] = runes[letter] - 32
		}
	}
	return string(runes)
}

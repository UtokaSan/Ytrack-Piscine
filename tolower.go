package piscine

func ToLower(s string) string {
	runes := []rune(s)

	for letter := range runes {
		if runes[letter] >= 'A' && runes[letter] <= 'Z' {
			runes[letter] = runes[letter] + 32
		}
	}
	return string(runes)
}

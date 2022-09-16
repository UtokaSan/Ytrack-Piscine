package piscine

func LastRune(s string) rune {
	length := len(s) - 1
	return []rune(s)[length]
}

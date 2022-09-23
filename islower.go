package piscine

func IsLower(s string) bool {
	for _, letter := range s {
		if letter > 122 || letter < 96 {
			return false
		}
	}
	return true
}

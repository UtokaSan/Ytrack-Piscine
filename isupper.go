package piscine

func IsUpper(s string) bool {
	for _, letter := range s {
		if letter > 90 || letter < 64 {
			return false
		}
	}
	return true
}

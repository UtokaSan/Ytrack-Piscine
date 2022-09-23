package piscine

func StrRev(s string) string {
	l := []rune(s)

	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
	return string(l)

}

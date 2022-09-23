package piscine

func ConcatParams(args []string) string {
	str := ""

	for i := 0; i <= len(args)-1; i++ {
		if i == len(args)-1 {
			str += args[i]
		} else {
			str += args[i] + "\n"
		}
	}
	return str
}

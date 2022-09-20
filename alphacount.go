package main

func AlphaCount(s string) int {
	count := 0

	for _, letter := range s {
		for alpha := 'a'; alpha <= 'z'; alpha++ {
			if letter == alpha {
				count++
			}
		}
		for alphaUpper := 'A'; alphaUpper <= 'Z'; alphaUpper++ {
			if letter == alphaUpper {
				count++
			}
		}
	}
	return count
}

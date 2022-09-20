package main

func IterativeFactorial(nb int) int {
	if nb > 20 || nb < 0 {
		return 0
	}

	if nb == 0 {
		return 1
	}
	stock := 1
	for i := 1; i <= nb; i++ {
		stock *= i
	}
	return stock
}

package piscine

func RecursivePower(nb int, power int) int {
	if nb == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}
	return nb * RecursivePower(nb, power-1)
}

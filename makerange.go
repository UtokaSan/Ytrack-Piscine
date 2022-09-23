package piscine

func MakeRange(min, max int) []int {
	size := max - min
	if min >= max {
		return nil
	}
	tab := make([]int, size)
	for i := 0; i < size; i++ {
		tab[i] = min + i
	}
	return tab
}

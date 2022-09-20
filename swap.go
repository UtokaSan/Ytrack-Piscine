package main

func swap(a *int, b *int) {
	temp := *a
	temp1 := *b
	*a = temp1
	*b = temp
}

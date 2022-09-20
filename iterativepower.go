package main

import "fmt"

func IterativePower(nb int, power int) int {

	if nb == 0 {
		return 0
	}

	stock := 1
	for i := 0; i < power; i++ {
		stock = stock * nb
		fmt.Println(i)
	}
	return stock
}

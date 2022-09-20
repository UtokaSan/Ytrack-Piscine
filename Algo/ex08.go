package main

import "fmt"

func main() {
	var maxi int = 0
	var n int = 0
	var indexMaxi int = 0

	for i := 1; i <= 20; i++ {
		fmt.Println("Entrer un nombre :")
		fmt.Scanf("%s", n)

		if n > maxi {
			maxi = n
			indexMaxi = i
		}
	}
	fmt.Println("Le Nombre le plus grand était : %i", n)
	fmt.Println("En position : %i", indexMaxi)
}

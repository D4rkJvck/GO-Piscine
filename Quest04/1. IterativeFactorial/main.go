package main

import "fmt"

func IterativeFactorial(nb int) int {
	//Conditions de Sortie
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	//Boucle à contre-sens: Factorielle
	for i := nb - 1; i > 0; i-- {
		nb = nb * i
	}
	//Résultat
	return nb
}

// TEST
func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}

package main

import "fmt"

func RecursivePower(nb int, power int) int {
	//Conditions de Sortie
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	//Factorisation avec Auto-Appel de la Fonction: Boucle autour de la Décrémentation de la Puissance
	return nb * RecursivePower(nb, power-1)
}

// TEST
func main() {
	fmt.Println(RecursivePower(4, 3))
}

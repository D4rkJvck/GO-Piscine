package main

import "fmt"

func IterativePower(nb int, power int) int {
	//Outil
	result := 1
	//Condition de Sortie
	if power < 0 {
		return 0
	}
	//Boucle: Puissance = Nombre d'Incrémentations (à partir de 1)
	for i := 1; i <= power; i++ {
		result = result * nb
	}
	//Résultat
	return result
}

// TEST
func main() {
	fmt.Println(IterativePower(4, 3))
}

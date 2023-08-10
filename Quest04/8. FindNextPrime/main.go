package main

import "fmt"

func FindNextPrime(nb int) int {
	//Condition de Sortie
	if nb <= 1 {
		return 2
	}
	//Boucle de Vérification liée à la Racine Carrée
	for i := 2; i*i <= nb; i++ {
		//Condition de Modularité
		if nb%i == 0 {
			//Incrémentation pour atteindre le plus proche nombre premier
			return FindNextPrime(nb + 1)
		}
	}
	return nb
}

// TEST
func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}

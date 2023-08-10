package main

import "fmt"

func IsPrime(nb int) bool {
	//Condition de Sortie liée à la Primalité
	if nb <= 1 {
		return false
	}
	//Boucle de Vérification liée à la Racine Carrée
	for i := 2; i*i <= nb; i++ {
		//Condition de Modularité
		if nb%i == 0 {
			return false
		}
	}
	return true
}

// TEST
func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}

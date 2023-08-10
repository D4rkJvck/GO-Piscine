package main

import "fmt"

func Sqrt(nb int) int {
	//Conditions de Sortie liées à la Racine Carrée
	if nb == 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	//Boucle de Vérification de la Racine Carrée
	for i := 1; i < nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}

// TEST
func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}

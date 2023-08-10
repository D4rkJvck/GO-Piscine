package main

import "fmt"

func Index(s string, toFind string) int {
	//Boucle où l'index s'incrémente tant qu'elle reste inférieure à la différence de taille entre les 2 chaînes
	for i := 0; i < len(s)-len(toFind); i++ {
		//Condition d'Arrêt
		if s[i:len(toFind)+i] == toFind {
			return i
		}
	}
	return -1
}

// TEST
func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}

package main

import "github.com/01-edu/z01"

func main() {
	//Boucle d'Affichage: Rune par Rune
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	//Retour à la ligne
	z01.PrintRune('\n') //===> z01.PrintRune(10)
}

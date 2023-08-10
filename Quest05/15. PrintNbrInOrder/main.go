package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var tab []rune //---------------------------------> Tableau de runes pour collecter les chiffres
	if n <= 0 {
		z01.PrintRune('0') //------------------------------------> Pas besoin d'ordonner en cas de 0
	}
	if n >= 1 { //--------------------------------------------------------------> Condition d'entrée
		for n != 0 {
			tab = append(tab, rune('0'+n%10)) //-----------------------------> Collecte des chiffres
			n /= 10                           //----------------> Elimination des chiffres collectés
		}
		for _, v := range tab { //---------------------------------> Boucles imbriquées comparatives
			for _, w := range tab {
				if v > w { //---------------------------------------------------------> Condition de
					v, w = w, v //-----------------------------------------------------> Permutation
				}
			}
		}
		for _, v := range tab { //-----------------------------------> Boucle d'Affichage du Tableau
			z01.PrintRune(v)
		}
	}
}

func main() { //------------------------------------------------------------------------------> TEST

	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}

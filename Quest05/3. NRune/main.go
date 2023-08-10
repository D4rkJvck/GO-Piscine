package main

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	//Transformation de la Chaîne en Tableau de Runes
	i := []rune(s)
	//Condition de Délimitation
	if n < 1 || n > len(i) {
		return 0
	}
	//Retour de la N-ième Rune ôtée de un pour l'indice
	return rune(i[n-1])
}

// TEST
func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}

package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	//Transformation de la chaîne en tableau de runes
	i := []rune(s)
	return i[0]
}

// TEST
func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}

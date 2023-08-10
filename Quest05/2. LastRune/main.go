package main

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	//Transformation de la chaîne et retour simultanné
	return rune(s[len(s)-1])
}

// TEST
func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}

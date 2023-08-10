package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := []rune(os.Args[0])       //-------------------------------------------------------------------------> Transformation des Arguments en Tableau de Rune
	for i := 2; i < len(arg); i++ { //-----------------------------------------------------------------> Analyse du Tableau en commençant par le deuxième indice
		z01.PrintRune(rune(arg[i]))
	}
	z01.PrintRune(10)
}

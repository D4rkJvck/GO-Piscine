package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]                   //-----------------------------------------------------------------------------> Variable de Délimitation des Arguments
	for i := len(arg) - 1; i >= 0; i-- { //-----------------------------------------------------------> Analyse du Fichier Argument par Argument en sens inverse
		for _, v := range arg[i] { //-------------------------------------------------------------------------------> Conversion des Arguments inversés en Runes
			z01.PrintRune(v)
		}
		z01.PrintRune(10)
	}
}

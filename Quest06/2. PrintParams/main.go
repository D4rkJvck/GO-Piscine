package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]      //------------------------------------------------------> Variable de Délimitation des Arguments: le premier étant le nom du fichier
	for _, v := range arg { //----------------------------------------------------------------------------------------> Analyse du Fichier Argument par Argument
		for _, r := range v { //---------------------------------------------------------------------------------------------> Conversion des Arguments en Runes
			z01.PrintRune(r)
		}
		z01.PrintRune(10)
	}
}

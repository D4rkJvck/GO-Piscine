package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Affichage d'une Chaîne
func printStr(s string) {
	for _, r := range s { //--------------------------------> String Scan
		z01.PrintRune(r) //---------------------------------> Printing
	}
	z01.PrintRune('\n') //-------------------------------------> New Line
}

// Vérification de Parité
func isEven(nbr int) bool { //----------------------------------> Even Checking
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

// Appel
func main() {
	//Variable de Délimitation des Arguments
	args := os.Args[1:]
	if isEven(len(args)) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}

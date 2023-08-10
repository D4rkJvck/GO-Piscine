package main

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	//Condition
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	//Retour à la ligne
	z01.PrintRune('\n') //===> z01.PrintRune(10)
}

// TEST
func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}

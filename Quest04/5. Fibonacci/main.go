package main

import "fmt"

func Fibonacci(index int) int {
	//Conditions de Sortie
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	/*Double Auto-Appel de la Focntion:
	Boucles autour de la Décrémentation du nombre
	suivant la formule de la Suite de FIBONACCI*/
	return Fibonacci(index-1) + Fibonacci(index-2)
}

// TEST
func main() {
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}

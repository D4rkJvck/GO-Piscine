package main

import "fmt"

func IsNumeric(s string) bool {
	//Analyse de la Chaîne Rune par Rune
	for _, v := range s {
		//Condition de Délimitation: Nombre
		if !(v >= '0' && v <= '9') {
			return false
		}
	}
	return true
}

// TEST
func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}

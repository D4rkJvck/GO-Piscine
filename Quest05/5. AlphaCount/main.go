package main

import "fmt"

func AlphaCount(s string) int {
	//Variable de Décompte
	c := 0
	//Analyse de la Chaîne Rune par Rune
	for _, v := range s {
		//Condition de Délimitation pour le Décompte
		if v >= 65 && v <= 90 || v >= 97 && v <= 122 {
			c++
		}
	}
	//Résultat
	return c
}

// TEST
func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}

package main

import "fmt"

func ToUpper(s string) string {
	//Transformation de la Chaîne en Tableau de Runes
	r := []rune(s)
	//Analyse de la Chaîne Rune par Rune
	for i := 0; i < len(s); i++ {
		//Condition de Délimitation des Runes
		if r[i] >= 97 && r[i] <= 122 {
			//Opération de Remplacement: Minuscule à Majuscule
			r[i] = r[i] - 32
		}
	}
	//Retour de la Transformation du Tableau de Rune en Chaîne
	return string(r)
}

// TEST
func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}

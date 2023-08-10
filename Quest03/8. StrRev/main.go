package main

import "fmt"

func StrRev(s string) string {
	// Outil
	var inverse string
	//Analyse de la chaîne à contre-sens
	for i := len(s) - 1; i >= 0; i-- {
		//Formule d'Inversion
		inverse = inverse + string(s[i])
	}
	//Résultat
	return inverse
}

// TEST
func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

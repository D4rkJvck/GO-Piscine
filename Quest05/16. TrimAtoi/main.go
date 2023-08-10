package main

import "fmt"

func TrimAtoi(s string) int {
	var atoi int //-----------------------------------------------------------------> Particule atoi
	neg := 1
	for _, v := range s { //-----------------------------------> Parcours de la Chaîne Rune par Rune
		if v >= '0' && v <= '9' { //---------------------------------------> Condition de Conversion
			atoi = atoi*10 + int(v-'0') //------------------> Formule de Conversion (A)scii-TO-(I)nt
		}
		if v == '-' && atoi == 0 { //----------------------------------------> Condition de Négation
			neg = -1
		}
	}
	return neg * atoi //------------------------------------------------> Résultat: avec la Négation
}

func main() { //------------------------------------------------------------------------------> TEST

	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}

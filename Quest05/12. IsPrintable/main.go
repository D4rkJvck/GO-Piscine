package main

import "fmt"

func IsPrintable(s string) bool {
	//Analyse de la Chaîne Rune par Rune
	for _, r := range s {
		//Condition de Délimitation: ASCII
		if !(r >= 32 && r <= 127) {
			return false
		}
	}
	return true
}

// TEST
func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))

}

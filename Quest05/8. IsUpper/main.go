package main

import "fmt"

func IsUpper(s string) bool {
	//Analyse de la chaîne Rune par Rune
	for i := 0; i < len(s); i++ {
		//Condition de Délimitation: Majuscule
		if !(rune(s[i]) >= 65 && rune(s[i]) <= 90) {
			return false
		}
	}
	return true
}

// TEST
func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))

}

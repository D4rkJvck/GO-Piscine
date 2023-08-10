package main

import "fmt"

func IsLower(s string) bool {
	//Analyse de la Chaîne Rune par Rune
	for i := 0; i < len(s); i++ {
		//Condition de Délimitation: Minuscule
		if !(rune(s[i]) >= 77 && rune(s[i]) <= 122) {
			return false
		}
	}
	return true
}

// TEST
func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}

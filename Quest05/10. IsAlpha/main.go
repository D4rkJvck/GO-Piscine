package main

import "fmt"

func IsAlpha(s string) bool {
	//Analyse de la Chaîne Rune par Rune
	for i := 0; i < len(s); i++ {
		//Condition de Délimitation: Alphanumérique
		if !((rune(s[i]) >= 'A' && rune(s[i]) <= 'Z') || (rune(s[i]) >= 'a' && rune(s[i]) <= 'z') || (rune(s[i]) >= '0' && rune(s[i]) <= '9')) {
			return false
		}
	}
	return true
}

// TEST
func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))

}

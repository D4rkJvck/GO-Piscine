package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	//Scanning String one Rune after another
	for i := 0; i < len(s); i++ { //================================================> for _, valeur:= range s
		z01.PrintRune(rune(s[i])) //====================================================> z01.PrintRune(valeur)
	}
}

func main() { //-------------------------------------------------------------------------> TEST

	PrintStr("Hello World!")
}

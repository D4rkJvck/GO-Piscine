package main

import "fmt"

func BasicAtoi(s string) int {
	var atoi int               //------------------------------------> Atoi Particle
	for _, valeur := range s { //----------------------------------> String Scan
		if valeur >= '0' && valeur <= '9' { //------------------------------------------> Digits Filter
			atoi = atoi*10 + int(valeur-'0') //------------------------> "(A)scii (TO) (I)nt" Formula
		}
	}
	return atoi //------------------------------------------------------------------------------> Result
}

func main() { //----------------------------------------------------------------------------> TEST
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

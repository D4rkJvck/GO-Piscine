package main

import (
	"github.com/01-edu/z01"
)

// Structure
type point struct {
	x int
	y int
}

// Pointage
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

// Affichage
func PrintStr(s string) {
	for _, value := range s {
		z01.PrintRune(value)
	}
}

// Conversion
func Itoa(n int) string {
	var itoa string //----------------------------> Variable qui va stocker les chiffres
	sign := ""      //-------------------------------------------> Particule de négation
	if n == 0 {
		return "0" //---------------------------------------------> Condition de sortie
	}
	if n < 0 {
		n *= -1    //------------------------------> Rendre n positif pour le manipuler
		sign = "-" //-----------------------------> Déclencher la particule de négation
	}
	for n != 0 { //-------------------------------> Boucle de Séparation des chiffres
		itoa = string(n%10+'0') + itoa //----------------------> Récupérer le dernier chiffre
		n /= 10                        //------------> Eliminer le chiffre après récupération
	}
	return sign + itoa //-------------------------------> Concaténer le signe au résultat
}

// Appel
func main() {
	// Variable de structure
	points := &point{}
	// Pointage
	setPoint(points)
	// Affichage
	PrintStr("x = ")
	PrintStr(Itoa(points.x))
	PrintStr(", y = ")
	PrintStr(Itoa(points.y))
	z01.PrintRune(10)
}

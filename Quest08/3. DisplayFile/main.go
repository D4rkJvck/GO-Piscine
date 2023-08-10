package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Variable de Délimitation des Arguments
	arg := os.Args[1:]
	//Conditions d'Affichage
	if len(arg) < 1 {
		fmt.Println("File name missing")
	}
	if len(arg) > 1 {
		fmt.Println("Too many arguments")
	}
	if len(arg) == 1 {
		//Ouverture et Lecture du Fichier
		file, _ := ioutil.ReadFile(os.Args[0])
		//Variable de Conversion (en Chaîne)
		str := string(file)
		//Affichage de la Chaîne convertie
		fmt.Println(str)
	}
}
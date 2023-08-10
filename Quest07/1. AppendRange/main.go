package main

import "fmt"

func AppendRange(min, max int) []int {
	//Condition de Sortie
	if min >= max {
		return nil
	}
	//Variable de Récupération: Segment
	var Slice []int
	//Analyse des Possibilités entre le min et max
	for i := min; i < max; i++ {
		Slice = append(Slice, i)
	}
	//Retour du Segment une fois rempli
	return Slice
}

//TEST
func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}

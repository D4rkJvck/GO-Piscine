package main

func MakeRange(min, max int) []int {
	//Condition de Sortie
	if min >= max {
		return nil
	}
	//Création de Tableau
	tab := make([]int, max-min)
	//Analyse du Tableau en utilisant l'Index
	for i := range tab {
		//Ajout de l'Index de chaque tour au minimum
		tab[i] = min + i
	}
	//Retour du Tableau
	return tab
}

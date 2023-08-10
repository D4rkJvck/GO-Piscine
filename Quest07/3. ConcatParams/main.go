package main

func ConcatParams(args []string) string {
	//Création d'une variable sans valeur
	var concat string
	//Analyse du Tableau de Chaînes donné en utilisant l'index et la valeur
	for i, v := range args {
		//Ajout successif de la valeur à la variable
		concat = concat + v
		//Condition dans une boucle: Retour à la ligne tant que l'index est inférieur au dernier élément
		if i < len(args)-1 {
			concat = concat + "\n"
		}
	}
	//Résultat
	return concat
}
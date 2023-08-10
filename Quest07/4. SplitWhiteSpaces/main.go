package main

func SplitWhiteSpaces(s string) []string {
	//Tableau de Chaîne initialement vide
	var mots []string
	//Variable de Contenu à ajouter dans le tableau
	var mot string
	//Analyse de la Chaîne donnée
	for _, lettre := range s {
		//Condition d'Ajout de mot dans mots
		if lettre == ' ' || lettre == '\t' || lettre == '\n' {
			//A Condition que le mot existe
			if mot != "" {
				mots = append(mots, mot)
				//Réinitialisation de mot pour le prochain tour
				mot = ""
			}
			//Condition d'Ajout de mot non respectée
		} else {
			mot = mot + string(lettre)
		}
	}
	//Condition d'existance du mot identique à celui à l'intérieur de la boucle
	if mot != "" {
		mots = append(mots, mot)
	}
	//Résultat
	return mots
}

package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	//===> Boucles imbriquées
	for i := '0'; i <= '9'; i++ {
		//dans la 1e Boucle
		for j := i + 1; j <= '9'; j++ {
			//dans la 2e Boucle
			for k := j + 1; k <= '9'; k++ {
				//3e Boucle
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				//Condition d'Arrêt
				if i < '7' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	//Retour à la ligne
	z01.PrintRune('\n') //===> z01.PrintRune(10)
}

// TEST
func main() {
	PrintComb()
}

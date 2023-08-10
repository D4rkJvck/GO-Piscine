package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	time := *a     //------------------------------> Storing *a's value in a time variable
	*a = *a / *b   //--------------------------------------> Re-assigning *a an operation (value change)
	*b = time % *b //------------------> Re-assigning *b an operation using the time variable's value (copy of *a)
}

func main() { //-----------------------------------------------------------------------------------> TEST

	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

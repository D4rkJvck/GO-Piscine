package main

import "fmt"

func DivMod(a int, b int, div *int, mod *int) {
	//Using pointers for operations
	*div = a / b
	*mod = a % b
}

func main() { //--------------------------------------------------------------------------------------> TEST
	a := 13
	b := 2
	var div int
	var mod int
	DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}

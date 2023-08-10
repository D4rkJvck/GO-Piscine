package main

import "fmt"

func Swap(a *int, b *int) {
	temporelle := *a //------------------------------> Storing *a's value in a time variable
	*a = *b          //--------------------------------------------------------> Re-assigning *a the value of b*
	*b = temporelle  //-------------------> Re-assigning *b the time variable's value (which is the value of *a)
}
func main() { //-------------------------------------------------------------------------> TEST
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

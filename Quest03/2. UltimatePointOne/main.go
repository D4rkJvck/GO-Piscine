package main

import "fmt"

func UltimatePointOne(n ***int) {
	***n = 1 //-----------------------------------------------------------> Pointer to a pointer to another one
}

func main() { //----------------------------------------------------------------------> TEST
	a := 0
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
}

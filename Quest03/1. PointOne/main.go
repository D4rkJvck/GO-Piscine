package main

import "fmt"

func PointOne(n *int) {
	*n = 1 //-------------------------------------> Pointing re-assignation
}

func main() { //------------------------------------------------------------------------------------------> TEST

	n := 0
	PointOne(&n)
	fmt.Println(n)
}

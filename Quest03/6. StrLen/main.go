package main

import "fmt"

func StrLen(s string) int {
	counter := 0  //---------------------------------------------> counter
	for range s { //-------------------------------------------------------------------> String Scan
		counter++ //-------------------------------------------------------> Counter till the end of the String
	}
	return counter //----------------------------------------------> Result
}

func main() { //-----------------------------------------------------------> TEST
	l := StrLen("Hello World!")
	fmt.Println(l)
}

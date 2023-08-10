package main

import "fmt"

func Concat(str1 string, str2 string) string {
	//Addition de 2 chaînes sans espace
	return str1 + str2
}

// TEST
func main() {
	fmt.Println(Concat("Hello!", " How are you?"))

}

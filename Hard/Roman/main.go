package main

import (
	"os"
	"fmt"
	"strconv"
)

func Roman(n int) string {

	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	if n > 0 && n < 4000 {
		return thousands[n/1000] + hundreds[(n%1000)/100] + tens[(n%100)/10] + ones[n%10]
	} else {
		return "ERROR: cannot convert to roman digit"
	}
}

func Roman2(n int) string {

	thousand := []string{"", "M", "M+M", "M+M+M"}
	hundred := []string{"", "C", "C+C", "C+C+C", "(D-C)", "D", "D+C", "D+C+C", "D+C+C+C", "(M-C)"}
	ten := []string{"", "X", "X+X", "X+X+X", "(L-C)", "L", "L+X", "L+X+X", "L+X+X+X", "(C-X)"}
	one := []string{"", "I", "I+I", "I+I+I", "(V-I)", "V", "V+I", "V+I+I", "V+I+I+I", "(X-I)"}

	if n > 0 && n < 10 {
		return one[n%10]
	} else if n > 10 && n < 100 {
		return ten[(n%100)/10] + "+" + one[n%10]
	} else if n > 100 && n < 1000 {
		return hundred[(n%1000)/100] + "+" + ten[(n%100)/10] + "+" + one[n%10]
	} else if n > 1000 && n < 4000 {
		return thousand[n/1000] + "+" + hundred[(n%1000)/100] + "+" + ten[(n%100)/10] + "+" + one[n%10]
	} else {
		return "ERROR: cannot convert to roman digit"
	}
}

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println(Roman(num))
	fmt.Println(Roman2(num))
}
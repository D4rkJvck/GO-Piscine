package main

import "github.com/01-edu/z01"

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

func PrintNbr(n int) {
	var tab []rune
	if n <= 0 {
		z01.PrintRune('0')
	}
	if n >= 1 {
		for n != 0 {
			tab = append(tab, rune('0'+n%10))
			n /= 10
		}
	}
	for _, v: range tab {
		z01.PrintRune(v)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}
s
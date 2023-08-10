package main

import "os"

func Atoi(s string) int {
	var atoi int
	signa := 1
	for i := 0; i < len(s); i++ {
		if s[0] == '-' {
			s = s[1:]
			signa *= -1

		} else if s[0] == '+' {
			s = s[1:]
		}
		if !(s[i] >= '0' && s[i] <= '9') {
			return 0
		} else {
			atoi = atoi*10 + int(s[i]-'0')
		}
	}
	return signa * atoi
}

func Itoa(n int) string {
	var s string
	signi := ""
	if n == 0 {
		return "0"
	}
	if n < 0 {
		n *= -1
		signi = "-"
	}
	for n != 0 {
		s = string(n%10+'0') + s
		n /= 10
	}
	return signi + s
}

// func IsNumeric(s string) bool {
// 	for _, n := range s {
// 		if !(n >= '0' && n <= '9') {
// 			return false
// 		}
// 	}
// 	return true
// }

func IsOperator(s string) bool {
	if !(s == "+" || s == "-" || s == "*" || s == "/" || s == "%") {
		return false
	}
	return true
}

func IsOverflow(n int) bool {
	if !(n < -9223372036854775807 && n > 9223372036854775807) {
		return true
	}
	return false
}

func Add(a, b int) int {
	return a + b
}
func Sub(a, b int) int {
	return a - b
}
func Multi(a, b int) int {
	return a * b
}
func Div(a, b int) int {
	return a / b
}
func Mod(a, b int) int {
	return a % b
}

func main() {
	// arg := os.Args[1:]
	if len(os.Args) > 4 {
		return
	}
	// if IsNumeric(os.Args[1]) && IsNumeric(os.Args[3]) {
	v1 := Atoi(os.Args[1])
	v2 := Atoi(os.Args[3])
	op := os.Args[2]
	if !(IsOverflow(v1) && IsOverflow(v2)) {
		if IsOperator(op) {
			if op == "+" {
				os.Stdout.WriteString(Itoa(Add(v1, v2)))
				os.Stdout.WriteString("\n")
			} else if op == "-" {
				os.Stdout.WriteString(Itoa(Sub(v1, v2)))
				os.Stdout.WriteString("\n")
			} else if op == "*" {
				os.Stdout.WriteString(Itoa(Multi(v1, v2)))
				os.Stdout.WriteString("\n")
			} else if op == "/" {
				if v2 == 0 {
					os.Stdout.WriteString("No division by 0")
					os.Stdout.WriteString("\n")
				} else {
					os.Stdout.WriteString(Itoa(Div(v1, v2)))
					os.Stdout.WriteString("\n")
				}
			} else if op == "%" {
				if v2 == 0 {
					os.Stdout.WriteString("No modulo by 0")
					os.Stdout.WriteString("\n")
				} else {
					os.Stdout.WriteString(Itoa(Mod(v1, v2)))
					os.Stdout.WriteString("\n")
				}
			}
		}
	}
	// }
}

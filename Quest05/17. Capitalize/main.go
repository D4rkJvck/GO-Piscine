package main

import "fmt"

func IsUpperCase(r rune) bool { //---------------------------------------------> Checking: UpperCase
	if r >= 'A' && r <= 'Z' {
		return true
	} else {
		return false
	}
}
func IsLowerCase(r rune) bool { //---------------------------------------------> Checking: LowerCase
	if r >= 'a' && r <= 'z' {
		return true
	} else {
		return false
	}
}
func IsNumber(r rune) bool { //--------------------------------------------------> Checking: Numbers
	if r >= '0' && r <= '9' {
		return true
	} else {
		return false
	}
}
func IsAlphaNum(r rune) bool { //-------------------------------------> Checking: Alphanumeric runes
	if IsUpperCase(r) || IsLowerCase(r) || IsNumber(r) {
		return true
	} else {
		return false
	}
}
func Capitalize(s string) string {
	str := []rune(s)         //------------------------------------------> Runes Table of the String
	if IsLowerCase(str[0]) { //--------------------------------> First String's index Capitalization
		str[0] -= 32
	}
	for i := 1; i < len(str); i++ { //---------------------------> String Scan Starting with index 1
		if IsUpperCase(str[i]) { //--------------------------------------> LowerCase for all Letters
			str[i] += 32
		}
		if !(IsAlphaNum(str[i-1])) && IsLowerCase(str[i]) { //---------------------> Words Condition
			str[i] -= 32
		}
	}
	return string(str) //---------------------------------> Result: String version of the Rune Table
}
func main() { //------------------------------------------------------------------------------> TEST

	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

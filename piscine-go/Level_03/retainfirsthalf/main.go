package main

import (
	"fmt"
)

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}

func RetainFirstHalf(str string) string {
	if str == "" {
		return ""
	}
	word := ""
	for i, char := range str {
		if i == 0 {
			word += string(char)
		} else if isUpper(char) {
			break
		} else {
			word += string(char)
		}
	}
	return word
}

func isUpper(char rune) bool {
	if char >= 'A' && char <= 'Z' {
		return true
	}
	return false
}

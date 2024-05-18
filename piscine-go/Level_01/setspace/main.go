package main

import (
	"fmt"
)

func main() {
	fmt.Println(SetSpace("HelloWorld"))
	fmt.Println(SetSpace("HelloWorld12"))
	fmt.Println(SetSpace("Hello World"))
	fmt.Println(SetSpace(""))
	fmt.Println(SetSpace("LoremIpsumWord"))
}

func SetSpace(s string) string {
	word := ""
	for i, char := range s {
		if i == 0 {
			word += string(char)
			continue
		} else if char >= 'a' && char <= 'z' {
			word += string(char)
		} else if char >= 'A' && char <= 'Z' {
			word += " " + string(char)
		} else {
			return "Error"
		}
	}
	return word
}

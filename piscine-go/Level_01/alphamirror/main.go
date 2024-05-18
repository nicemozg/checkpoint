package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]
	finalWord := ""
	for _, char := range args {
		if char >= 'A' && char <= 'Z' {
			char = 'Z' - char + 'A'
			finalWord += string(char)
		} else if char >= 'a' && char <= 'z' {
			char = 'z' - char + 'a'
			finalWord += string(char)
		} else {
			finalWord += string(char)
		}
	}
	fmt.Println(finalWord)
}

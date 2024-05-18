package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}

	str1 := args[0]
	str2 := args[1]

	charMap := make(map[rune]bool)
	for _, char := range str2 {
		charMap[char] = true
	}

	var result string
	for _, char := range str1 {
		if charMap[char] {
			result += string(char)
			charMap[char] = false
		}
	}
	fmt.Println(result)
}

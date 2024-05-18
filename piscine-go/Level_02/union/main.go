package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println()
		return
	}
	word := args[0] + args[1]
	myMap := make(map[rune]bool)
	result := ""
	for _, char := range word {
		if _, exists := myMap[char]; !exists {
			myMap[char] = true
			result += string(char)
		}
	}
	os.Stdout.WriteString(result)
}

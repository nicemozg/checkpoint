package main

import "os"

func main() {
	args := os.Args[1:]
	word := args[0]
	firstSymbol := []byte(args[1])
	secondSymbol := []byte(args[2])
	if len(args) > 3 {
		return
	}
	newWord := ""

	for _, char := range word {
		if char == rune(firstSymbol[0]) {
			char = rune(secondSymbol[0])
		}
		newWord += string(char)
	}
	os.Stdout.WriteString(newWord)
}

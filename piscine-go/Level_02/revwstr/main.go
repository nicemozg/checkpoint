package main

import "os"

func main() {
	arg := os.Args[1]
	word := ""
	currentWord := ""
	for _, char := range arg {
		if char == ' ' {
			word = currentWord + " " + word
			currentWord = ""
		} else {
			currentWord += string(char)
		}
	}
	word = currentWord + word + "\n"
	os.Stdout.WriteString(word)
}

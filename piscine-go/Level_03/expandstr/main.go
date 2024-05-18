package main

import "os"

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		return
	}

	currentWord := ""
	array := []string{}
	for _, char := range arg[0] {
		if char == ' ' && currentWord != "" {
			array = append(array, currentWord)
			currentWord = ""
		} else if char != ' ' {
			currentWord += string(char)
		}
	}
	if currentWord != "" {
		array = append(array, currentWord)
		currentWord = ""
	}
	for i, word := range array {
		if i != len(array)-1 {
			currentWord += word + "   "
		} else {
			currentWord += word
		}
	}
	os.Stdout.WriteString(currentWord)
}

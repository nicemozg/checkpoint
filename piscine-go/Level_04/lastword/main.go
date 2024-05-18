package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}

	word := ""
	result := ""
	flag := false
	for _, char := range args[0] {
		if char == ' ' && flag {
			result = word
			word = ""
			flag = false
		} else if char != ' ' {
			word += string(char)
			flag = true
		}
	}
	if word != "" {
		os.Stdout.WriteString(word)
	} else {
		os.Stdout.WriteString(result)
	}
}

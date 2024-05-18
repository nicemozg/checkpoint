package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	} else if len(args) > 1 {
		os.Stdout.WriteString("\n")
		return
	}
	word := ""
	for _, char := range args[0] {
		if char != ' ' {
			word += string(char)
		} else if char == ' ' && word != "" {
			break
		}
	}
	os.Stdout.WriteString(word)
}

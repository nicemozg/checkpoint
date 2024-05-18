package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		word := ""
		lastWord := ""

		for _, char := range args[0] {
			if char == ' ' {
				if word != "" {
					lastWord = word
				}
				word = ""
			} else {
				word += string(char)
			}
		}
		if word != "" {
			lastWord = word
		}
		if lastWord != "" {
			os.Stdout.WriteString(lastWord)
			os.Stdout.WriteString("\n")
		}

	}

}

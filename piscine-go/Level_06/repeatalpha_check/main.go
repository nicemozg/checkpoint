package main

import "os"

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	if arg == "" {
		os.Stdout.WriteString("\n")
		return
	}
	word := ""
	count := 0
	for _, char := range arg {
		if char >= 'A' && char <= 'Z' {
			count = int(char-'A') + 1
		} else if char >= 'a' && char <= 'z' {
			count = int(char-'a') + 1
		} else {
			count = 1
		}
		for i := 0; i < count; i++ {
			word += string(char)
		}
	}
	os.Stdout.WriteString(word + "\n")
}

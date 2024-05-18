package main

import "os"

func main() {
	args := os.Args[1:]
	argFirst := args[0]
	argSecond := args[1]

	index := 0
	word := ""

	for _, char := range argFirst {
		for i := index; i < len(argSecond); i++ {
			if char == rune(argSecond[i]) {
				word += string(char)
				index = i + 1
				break
			}
		}
	}
	if argFirst == word {
		os.Stdout.WriteString(word)
	}
}

package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	word_01 := args[0]
	word_02 := args[1]

	index := 0
	result := ""
	for _, char := range word_01 {
		for i := index; i < len(word_02); i++ {
			if char == rune(word_02[i]) {
				result += string(char)
				index = i + 1
				break
			}
		}
	}
	if word_01 == result {
		os.Stdout.WriteString("1")
	} else {
		os.Stdout.WriteString("0")
	}
}

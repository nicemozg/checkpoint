package main

import "os"

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("\n")
		return
	}
	arg := os.Args[1]
	if len(arg) == 0 {
		os.Stdout.WriteString("\n")
		return
	}
	array := []string{}
	word := ""
	flag := false
	for _, char := range arg {
		if char == ' ' && flag {
			array = append(array, word)
			word = ""
			flag = false
		} else if char != ' ' {
			word += string(char)
			flag = true
		}
	}
	if word != "" {
		array = append(array, word)
	}
	word = ""
	for i := 1; i < len(array); i++ {
		word += array[i] + " "
	}
	word += array[0]
	os.Stdout.WriteString(word + "\n")
}

package main

import "os"

func main() {
	if len(os.Args) == 2 {
		args := os.Args[1]
		lastWord := ""

		for i := len(args) - 1; i >= 0; i-- {
			if args[i] != ' ' {
				lastWord = string(args[i]) + lastWord
			} else if lastWord != "" {
				break
			}
		}

		if lastWord != "" {
			println(lastWord)
		}
	}
}

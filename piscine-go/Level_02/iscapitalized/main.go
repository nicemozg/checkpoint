package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

func IsCapitalized(s string) bool {

	if len(s) == 0 {
		return false
	}
	flag := false
	for i, char := range s {
		if i == 0 && char >= 'a' && char <= 'z' {
			return false
		} else if char == ' ' {
			flag = true
		} else if flag {
			if char >= 'a' && char <= 'z' {
				return false
			} else {
				flag = false
			}
		}
	}
	return true
}

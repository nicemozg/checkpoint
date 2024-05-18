package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	} else if len(args) == 2 {
		num_01 := args[0]
		num_02 := args[1]
		division := 0
		if atoi(num_01) > atoi(num_02) {
			division = atoi(num_02)
		} else {
			division = atoi(num_01)
		}
		for i := division; i > 0; i-- {
			if atoi(num_01)%i == 0 && atoi(num_02)%i == 0 {
				fmt.Println(i)
				break
			}
		}
	} else {
		os.Stdout.WriteString("\n")
	}
}

func atoi(num string) int {
	number := 0
	for _, char := range num {
		number = number*10 + int(char-'0')
	}
	return number
}

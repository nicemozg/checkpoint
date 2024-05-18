package main

import (
	"fmt"
	"os"
)

func matchBrackets(exp string) bool {
	expArray := []rune(exp)
	var resultArray []rune
	index := -1
	for _, char := range expArray {
		if char == '(' || char == '[' || char == '{' {
			resultArray = append(resultArray, char)
			index++
		} else if char == ')' {
			if index < 0 || resultArray[index] != '(' {
				return false
			}
			resultArray = resultArray[:len(resultArray)-1]
			index--
		} else if char == '}' {
			if index < 0 || resultArray[index] != '{' {
				return false
			}
			resultArray = resultArray[:len(resultArray)-1]
			index--
		} else if char == ']' {
			if index < 0 || resultArray[index] != '[' {
				return false
			}
			resultArray = resultArray[:len(resultArray)-1]
			index--
		}
	}
	return len(resultArray) == 0
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println()
	} else {
		for _, v := range os.Args[1:] {
			if matchBrackets(v) {
				fmt.Println("OK")
			} else {
				fmt.Println("Error")
			}
		}
	}
}

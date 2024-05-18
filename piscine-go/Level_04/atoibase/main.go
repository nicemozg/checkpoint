package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}
	for _, char := range base {
		if char == '-' || char == '+' {
			return 0
		}
	}

	baseMap := make(map[rune]int)
	for i, char := range base {
		if _, exsist := baseMap[char]; exsist {
			return 0
		}
		baseMap[char] = i
	}

	result := 0
	baseLen := len(base)

	for _, char := range s {
		if _, exsist := baseMap[char]; !exsist {
			return 0
		}
		result = result*baseLen + baseMap[char]
	}
	return result
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	input, err := strconv.Atoi(args[0])
	if err != nil || input < 0 {
		fmt.Println("ERROR" + "\n")
		return
	}
	fmt.Println(decimalToHexadecimal(input))
}

func decimalToHexadecimal(n int) string {
	hexChars := "0123456789abcdef"
	var result string

	for n > 0 {
		remainder := n % 16
		result = string(hexChars[remainder]) + result
		n = n / 16
	}

	if result == "" {
		result = "0"
	}

	return result
}

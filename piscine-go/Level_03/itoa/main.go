package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	sign := ""
	if n < 0 {
		sign += "-"
		n *= -1
	}
	for n > 0 {
		digit := n % 10
		result = string(rune(digit+'0')) + result
		n /= 10
	}
	return sign + result
}

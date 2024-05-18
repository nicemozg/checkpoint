package main

import "os"

func main() {
	args := os.Args[1:]
	count := len(args)
	os.Stdout.WriteString(itoa(count))
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		digit := n % 10
		result = string(rune(digit+'0')) + result
		n /= 10
	}
	return result
}

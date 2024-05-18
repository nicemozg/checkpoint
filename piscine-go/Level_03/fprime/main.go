package main

import (
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	num, err := strconv.Atoi(args[0])
	if err != nil || num < 2 {
		return
	} else {
		fprime(num)
	}
}

func fprime(num int) {
	division := 2
	result := ""
	for num > 1 {
		if num%division == 0 {
			result += itoa(division)
			num = num / division
			if num > 1 {
				result += "*"
			}
			division--
		}
		division++
	}
	os.Stdout.WriteString(result)
}

func itoa(num int) string {
	result := ""
	for num > 0 {
		digit := num % 10
		result = string(rune(digit+'0')) + result
		num = num / 10
	}
	return result
}

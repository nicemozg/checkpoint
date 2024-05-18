package main

import (
	"fmt"
)

func main() {
	fmt.Println(NotDecimal("0.1"))
	fmt.Println(NotDecimal("174.2"))
	fmt.Println(NotDecimal("0.1255"))
	fmt.Println(NotDecimal("1.20525856"))
	fmt.Println(NotDecimal("-0.0f00d00"))
	fmt.Println(NotDecimal(""))
	fmt.Println(NotDecimal("-19.525856"))
	fmt.Println(NotDecimal("1952"))
}

func NotDecimal(dec string) string {

	result := ""
	for _, char := range dec {
		if char >= 'A' && char <= 'Z' {
			return dec
		} else if char >= 'a' && char <= 'z' {
			return dec
		}
	}
	numNotDecimal := atoi(dec)
	result = itoa(numNotDecimal)
	return result
}

func atoi(dec string) int {
	num := 0
	sign := 1
	for i, char := range dec {
		if i == 0 && char == '-' {
			sign *= -1
		} else if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else if char == '.' {
			continue
		}
	}
	return num * sign
}

func itoa(num int) string {
	result := ""
	sign := ""
	if num < 0 {
		num *= -1
		sign = "-"
	}
	for num > 0 {
		digit := num % 10
		result = string(rune(digit+'0')) + result
		num /= 10
	}
	return sign + result
}

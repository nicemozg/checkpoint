package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println(0)
		return
	}
	result := 0
	num := atoi(args[0])
	for i := num; i > 1; i-- {
		if isPrime(i) {
			result += i
		}
	}
	fmt.Println(result)
}

func isPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func atoi(num string) int {
	result := 0
	for _, char := range num {
		result = result*10 + int(char-'0')
	}
	return result
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPowerOf2(n int) bool {
	return n != 0 && (n&(n-1)) == 0
}

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		num, err := strconv.Atoi(args[0])
		if err == nil {
			fmt.Println(isPowerOf2(num))
		}
	}
}

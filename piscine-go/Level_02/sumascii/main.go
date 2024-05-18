package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("0")
		return
	}
	arrByte := []byte(args[0])
	var result byte
	for _, byte := range arrByte {
		result += byte
	}
	fmt.Println(result)
}

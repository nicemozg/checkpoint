package main

import (
	"fmt"
	"os"
	"strconv"
)
import "strings"

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("Error" + "\n")
		return
	}
	op := strings.Split(os.Args[1], " ")
	op = deleteSpace(op)
	values := []int{}

	for _, char := range op {
		num, err := strconv.Atoi(char)
		if err == nil {
			values = append(values, num)
			continue
		}
		n := len(values)
		if n < 2 {
			fmt.Println("Error")
			return
		}
		switch char {
		case "+":
			values[n-2] += values[n-1]
			values = values[:n-1]
		case "-":
			values[n-2] -= values[n-1]
			values = values[:n-1]
		case "*":
			values[n-2] *= values[n-1]
			values = values[:n-1]
		case "/":
			values[n-2] /= values[n-1]
			values = values[:n-1]
		case "%":
			values[n-2] %= values[n-1]
			values = values[:n-1]
		default:
			fmt.Println("Error")
			return
		}
	}
	if len(values) == 1 {
		fmt.Println(values[0])
	} else {
		fmt.Println("Error")
	}
}

func deleteSpace(a []string) (b []string) {
	array := []string{}
	for _, char := range a {
		if char != "" {
			array = append(array, char)
		}
	}
	return array
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(SwapFirst([]int{1, 2, 3, 4}))
	fmt.Println(SwapFirst([]int{3, 4, 6}))
	fmt.Println(SwapFirst([]int{1}))
}

func SwapFirst(slice []int) []int {
	if len(slice) > 1 {
		slice[0], slice[1] = slice[1], slice[0]
		return slice
	}
	return slice
}

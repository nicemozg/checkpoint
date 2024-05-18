package main

import (
	"fmt"
)

func main() {
	fmt.Println(SwapLast([]int{1, 2, 3, 4}))
	fmt.Println(SwapLast([]int{3, 4, 5}))
	fmt.Println(SwapLast([]int{1}))
}

func SwapLast(slice []int) []int {
	if len(slice) > 1 {
		slice[len(slice)-1], slice[len(slice)-2] = slice[len(slice)-2], slice[len(slice)-1]
		return slice
	}
	return slice
}

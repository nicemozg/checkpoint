package main

import (
	"fmt"
)

func main() {
	fmt.Println(BeZero([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(BeZero([]int{}))
}

func BeZero(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}
	for i := 0; i < len(slice); i++ {
		slice[i] = 0
	}
	return slice
}

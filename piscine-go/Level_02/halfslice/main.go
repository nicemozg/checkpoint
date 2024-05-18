package main

import (
	"fmt"
)

func main() {
	fmt.Println(HalfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(HalfSlice([]int{1, 2, 3}))
}

func HalfSlice(slice []int) []int {
	length := 0
	if len(slice)%2 == 0 {
		length = len(slice) / 2
	} else {
		length = (len(slice) + 1) / 2
	}
	return slice[:length]
}

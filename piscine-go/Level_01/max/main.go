package main

import (
	"fmt"
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	max := a[0]
	for _, num := range a {
		if max < num {
			max = num
		}
	}
	return max
}

package main

import (
	"fmt"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	ascending, descending := true, true
	for i := 0; i < n-1; i++ {
		result := f(a[i], a[i+1])
		if result > 0 {
			ascending = false
		} else if result < 0 {
			descending = false
		}
	}
	return ascending || descending
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

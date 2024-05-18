package main

import (
	"fmt"
)

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}

func SortWordArr(a []string) {

	for i := 0; i < len(a); i++ {
		for j := i; j < len(a)-1; j++ {
			if (a[i]) > (a[j+1]) {
				a[i], a[j+1] = a[j+1], a[i]
			}
		}
	}
}

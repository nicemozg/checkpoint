package main

import "fmt"

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

func ConcatAlternate(slice1, slice2 []int) []int {

	slice1Length := len(slice1)
	slice2Length := len(slice2)
	if slice1Length == 0 {
		return slice2
	} else if slice2Length == 0 {
		return slice1
	}
	sum := slice1Length + slice2Length
	array := []int{}
	index1 := 0
	index2 := 0
	if slice1Length < slice2Length {
		for i := 0; i < sum; i++ {
			if i%2 == 0 {
				array = append(array, slice2[index2])
				index2++
			} else {
				if slice1Length < index1+1 {
					sum++
					continue
				} else {
					array = append(array, slice1[index1])
					index1++
				}
			}
		}
	} else {
		for i := 0; i < sum; i++ {
			if i%2 == 0 {
				if slice1Length < index1+1 {
					sum++
					continue
				} else {
					array = append(array, slice1[index1])
					index1++
				}
			} else {
				array = append(array, slice2[index2])
				index2++
			}
		}
	}
	return array
}

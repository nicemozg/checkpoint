package main

import (
	"fmt"
)

func main() {
	fmt.Println(ReverseSecondHalf("This is the 1st half This is the 2nd half"))
	fmt.Println(ReverseSecondHalf(""))
	fmt.Println(ReverseSecondHalf("Hello World"))
}

func ReverseSecondHalf(str string) string {
	if str == "" {
		return "Invalid Output"
	}
	secondHalf := ""
	flag := false
	for i, char := range str {
		if i == 0 {
			continue
		} else if flag {
			secondHalf += string(char)
		} else if char >= 'A' && char <= 'Z' {
			secondHalf += string(char)
			flag = true
		}
	}
	array := []byte(secondHalf)
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(secondHalf)-i-1] = array[len(secondHalf)-i-1], array[i]
	}
	return string(array)
}

package main

import (
	"os"
)

func main() {
	arg := os.Args[1]
	result := Rot13(arg)
	os.Stdout.WriteString(result)
}

func Rot13(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			if r > 'm' {
				r -= 13
			} else {
				r += 13
			}
		} else if r >= 'A' && r <= 'Z' {
			if r > 'M' {
				r -= 13
			} else {
				r += 13
			}
		}
		result += string(r)
	}
	return result
}

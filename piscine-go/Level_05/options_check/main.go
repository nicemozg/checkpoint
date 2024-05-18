package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println(Options(args...))
}

func Options(args ...string) string {
	if len(args) == 0 {
		return "options_check: abcdefghijklmnopqrstuvwxyz"
	}
	options := 0
	showHelp := false

	for _, arg := range args {
		if arg == "-" {
			return "Invalid Option"
		} else if arg[:2] == "-h" {
			showHelp = true
			break
		}
		if len(arg) > 1 && arg[0] == '-' {
			for _, char := range arg[1:] {
				if char >= 'a' && char <= 'z' {
					options |= 1 << (char - 'a')
				} else {
					return "Invalid Option"
				}
			}
		} else {
			return "Invalid Option"
		}
	}

	if showHelp {
		return "options_check: abcdefghijklmnopqrstuvwxyz"
	}

	binaryStr := fmt.Sprintf("%032b", options)
	result := ""
	for i, bit := range binaryStr {
		result += string(bit)
		if (i+1)%8 == 0 && i != 31 {
			result += " "
		}
	}

	return result
}

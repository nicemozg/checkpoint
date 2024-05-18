package main

import "github.com/01-edu/z01"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	for i := 0; i < 10; i++ {
		printHex(arr[i])
		if i%4 == 3 {
			z01.PrintRune('\n')
		} else if i < 9 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
	for i := 0; i < 10; i++ {
		if int(arr[i]) >= 32 && int(arr[i]) <= 126 {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func printHex(n byte) {
	hex := "0123456789abcdef"
	z01.PrintRune(rune(hex[n>>4]))
	z01.PrintRune(rune(hex[n&0x0f]))
}

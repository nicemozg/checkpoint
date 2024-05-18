package main

import "os"
import "fmt"

func main() {
	args := os.Args[1:]
	count := 0
	if len(args) > 0 {
		for _, word := range args {
			for range word {
				count++
			}
		}
	}
	fmt.Println(count)
}

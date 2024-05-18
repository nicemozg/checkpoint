package main

import "fmt"

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	word := []rune(s)
	newWord := ""

	for i := 0; i < len(s); i++ {
		if toUpper(word[i]) && toUpper(word[i+1]) {
			return s
		} else if i == 0 {
			newWord += string(word[i])
			continue
		} else {
			if toUpper(word[i]) {
				newWord += "_" + string(word[i])
			} else {
				newWord += string(word[i])
			}
		}
	}
	return newWord
}

func toUpper(char rune) bool {
	if char >= 'A' && char <= 'Z' {
		return true
	}
	return false
}

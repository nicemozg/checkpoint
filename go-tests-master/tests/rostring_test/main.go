package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"abc   ",
		"Let there     be light",
		"     AkjhZ zLKIJz , 23y",
		"",
	}
	args = append(args, random.StrSlice(chars.Words)...)

	for _, arg := range args {
		challenge.Program("rostring_check", arg)
	}
	challenge.Program("rostring_check")
	challenge.Program("rostring_check", "this", "is")
	challenge.Program("rostring_check", "not", "good", "for  you")
}

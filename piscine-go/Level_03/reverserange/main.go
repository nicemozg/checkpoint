package main

import (
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	num_01, err_01 := strconv.Atoi(args[0])
	num_02, err_02 := strconv.Atoi(args[1])
	if err_01 != nil {
		os.Stdout.WriteString("strconv.Atoi: parsing " + "\"" + args[0] + "\": invalid syntax$")
		return
	} else if err_02 != nil {
		os.Stdout.WriteString("strconv.Atoi: parsing " + "\"" + args[1] + "\": invalid syntax$")
		return
	}
	result := ""
	if num_01 < num_02 {
		for i := num_02; i >= num_01; i-- {
			if i == num_01 {
				result += itoa(i)
			} else {
				result += itoa(i) + " "
			}
		}
	} else {
		for i := num_02; i <= num_01; i++ {
			if i == num_01 {
				result += itoa(i)
			} else {
				result += itoa(i) + " "
			}
		}
	}

	os.Stdout.WriteString(result)
}

func itoa(num int) string {
	if num == 0 {
		return "0"
	}
	result := ""
	sighn := ""
	if num < 0 {
		num *= -1
		sighn = "-"
	}
	for num > 0 {
		digit := num % 10
		result += string(rune(digit + '0'))
		num /= 10
	}
	return sighn + result
}

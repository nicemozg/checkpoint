package main

import "os"

func main() {
	args := os.Args[1:]
	num_01, err_01 := atoi(args[0])
	num_02, err_02 := atoi(args[1])
	if !err_01 {
		os.Stdout.WriteString("strconv.Atoi: parsing " + "\"" + args[0] + "\": invalid syntax")
	} else if !err_02 {
		os.Stdout.WriteString("strconv.Atoi: parsing " + "\"" + args[1] + "\": invalid syntax")
	}

	result := ""
	for i := num_01; i <= num_02; i++ {
		if i == num_02 {
			result += itoa(i)
		} else {
			result += itoa(i) + " "
		}
	}
	os.Stdout.WriteString(result)
}

func atoi(str string) (int, bool) {
	result := 0
	sighn := 1
	for i, char := range str {
		if i == 0 && char == '-' {
			sighn *= -1
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		} else {
			return 0, false
		}
	}
	return result * sighn, true
}

func itoa(num int) string {
	if num == 0 {
		return "0"
	}
	result := ""
	sighn := ""
	if num < 0 {
		sighn += "-"
		num *= -1
	}
	for num > 0 {
		digit := num % 10
		result = string(rune(digit+'0')) + result
		num /= 10
	}
	return sighn + result
}

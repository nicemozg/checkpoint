package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := "999"
	num, _ := strconv.Atoi(arg)
	res, _ := ToRoman(num)
	romanString := RomanCalculate(res)
	os.Stdout.WriteString(romanString + "\n")
	os.Stdout.WriteString(res + "\n")

}

func ToRoman(n int) (string, error) {
	if n <= 0 || n >= 4000 {
		return "", fmt.Errorf("ERROR: cannot convert to roman digit")
	}

	var result string
	for _, numeral := range RomanNumerals {
		for n >= numeral.Value {
			result += numeral.Symbol
			n -= numeral.Value
		}
	}
	return result, nil
}

func RomanCalculate(roman string) string {
	result := ""
	for i := 0; i < len(roman)-1; i++ {
		if Symbol(string(roman[i])) >= Symbol(string(roman[i+1])) {
			if i == len(roman)-2 {
				result = result + string(roman[i]) + "+" + string(roman[i+1])
			} else {
				result = result + string(roman[i]) + "+"
			}
		} else {
			if i == len(roman)-2 {
				result = result + "(" + string(roman[i+1]) + "-" + string(roman[i]) + ")"
			} else {
				result = result + "(" + string(roman[i+1]) + "-" + string(roman[i]) + ")" + "+"
				i += 1
			}
		}
	}
	return result
}

func Symbol(symbol string) int {
	for _, numeral := range RomanNumerals {
		if numeral.Symbol == symbol {
			return numeral.Value
		}
	}
	return 0
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

// RomanNumerals represents a slice of Roman numerals.
var RomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

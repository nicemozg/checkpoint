package piscine

func ItoaBase(value, base int) string {
	if value == 0 {
		return "0"
	}
	if base < 2 || base > 16 {
		return ""
	}
	result := ""
	sign := ""

	str := "0123456789ABCDEF"
	strBase := str[:base]
	if value < 0 {
		sign = "-"
		for value < 0 {
			index := value % base * -1
			result = string(strBase[index]) + result
			value /= base
		}
	} else {
		for value > 0 {
			index := value % base
			result = string(strBase[index]) + result
			value /= base
		}
	}
	return sign + result
}

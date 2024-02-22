package utils

import "strconv"

func AddZeroBeforeNumber(number int) string {
	if number < 10 {
		return "0" + string(number)
	}
	return string(number)
}

func AddZeroBeforeNumberFromStr(number string) string {
	if len(number) == 0 {
		return "01"
	}
	// if number start with 0, remove it
	if number[0] == '0' {
		return number
	}
	n, _ := strconv.Atoi(number)
	if int(n) < 10 {
		return "0" + number
	}
	return string(number)
}

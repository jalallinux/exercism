package romannumerals

import "errors"

type romanNumeral struct {
	value   int
	numeral string
}

var romanNumerals = []romanNumeral{
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

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("input must be between 1 and 3999")
	}

	var result string

	for _, rn := range romanNumerals {
		for input >= rn.value {
			result += rn.numeral
			input -= rn.value
		}
	}

	return result, nil
}

package converter

import "strings"

var romanNumeralToIntMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func ConvertRomanNumeralToInteger(romanNumeralStr string) int {
	romanNumeralStr = strings.ToUpper(romanNumeralStr)

	if romanNumeralStr == "" {
		return -1
	} else {
		for i := 0; i < len(romanNumeralStr); i++ {
			if _, exists := romanNumeralToIntMap[romanNumeralStr[i]]; !exists {
				return -1
			}
		}

		total := 0
		lengthOfRomanNumeralStr := len(romanNumeralStr)
		for i := 0; i < lengthOfRomanNumeralStr; i++ {
			intValueOfSingleChar := romanNumeralToIntMap[romanNumeralStr[i]]

			if i < lengthOfRomanNumeralStr-1 && intValueOfSingleChar < romanNumeralToIntMap[romanNumeralStr[i+1]] {
				total -= intValueOfSingleChar
			} else {
				total += intValueOfSingleChar
			}
		}

		return total
	}
}

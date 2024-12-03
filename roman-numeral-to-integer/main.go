package main

import (
	"assignment1/roman-numeral-to-integer/converter"
	"fmt"
)

func main() {
	var romanNumeralStr string

	fmt.Println("Program to convert roman numeral string to integer")

	fmt.Println("Please enter a roman numeral string:")

	fmt.Scanln(&romanNumeralStr)

	intValueForRomanNumeralStr := converter.ConvertRomanNumeralToInteger(romanNumeralStr)

	if intValueForRomanNumeralStr == -1 {
		fmt.Println("Entered roman numeral string is invalid!")
	} else {
		fmt.Printf("Integer value of roman numeral string - %s is %d", romanNumeralStr, intValueForRomanNumeralStr)
	}
}

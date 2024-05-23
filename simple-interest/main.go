package main

import (
	"assignment1/simple-interest/calculator"
	"fmt"
)

func main() {
	var principalAmount, rate float64
	var timeInYears int

	fmt.Printf("Program to calculate the simple interest\n\n")

	fmt.Printf("Please enter the principal amount:\n")
	fmt.Scanln(&principalAmount)

	fmt.Printf("Please enter the rate:\n")
	fmt.Scanln(&rate)

	fmt.Printf("Please enter the time period in years:\n")
	fmt.Scanln(&timeInYears)

	var simpleInterest float64 = calculator.CalculateSimpleInterest(principalAmount, rate, timeInYears)

	fmt.Printf("\nSimple interest for the principal amount of %f with the rate of %f%% for %d years is: %.2f", principalAmount, rate, timeInYears, simpleInterest)
	// fmt.Printf("%.2f", simpleInterest)
}

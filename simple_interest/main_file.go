package main

import (
	"assignment1/simple_interest/calculator"
	"fmt"
)

func main() {
	principalAmount, rate, timeInYears := 500000.78965, 2.5, 4

	var simpleInterest float64 = calculator.CalculateSimpleInterest(principalAmount, rate, timeInYears)

	fmt.Print("Simple interest for the principal amount of ", principalAmount, " with the rate of ", rate, "% for ", timeInYears, " years is: ", simpleInterest)
	// fmt.Printf("%.2f", simpleInterest)
}

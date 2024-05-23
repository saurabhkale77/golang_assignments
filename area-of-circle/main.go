package main

import (
	"assignment1/area-of-circle/calculator"
	"fmt"
)

func main() {
	var radius float64

	fmt.Printf("Program to calculate the area of a circle\n\n")

	fmt.Printf("Please enter the radius of the circle:\n")
	fmt.Scanln(&radius)

	fmt.Printf("\nArea of circle with radius - %f is: %.2f", radius, calculator.CalculateAreaOfCircle(radius))
}

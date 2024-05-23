package main

import (
	"assignment1/area-of-circle/calculator"
	"fmt"
)

func main() {
	var radius float64 = 7.5

	fmt.Print("Area of circle with radius - ", radius, " is: ", calculator.CalculateAreaOfCircle(radius))
}

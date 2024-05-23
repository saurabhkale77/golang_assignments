package calculator

import "math"

func CalculateAreaOfCircle(radius float64) float64 {
	areaOfCircle := math.Pi * math.Pow(radius, 2)

	return math.Round(areaOfCircle*100) / 100
}

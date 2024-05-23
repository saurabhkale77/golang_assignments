package calculator

import "math"

func CalculateSimpleInterest(principalAmount, rate float64, timeInYears int) float64 {
	simpleInterest := principalAmount * rate * 0.01 * float64(timeInYears)

	return math.Round(simpleInterest*100) / 100
}

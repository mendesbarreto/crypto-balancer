package util

import "math"

func Round(number float64) int {
	return int(number + math.Copysign(0.5, number))
}

func Fixed(number float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(number*output)) / output
}

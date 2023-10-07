package darts

import "math"

func Score(x, y float64) int {
	a := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	if a <= 1 {
		return 10
	} else if a <= 5 {
		return 5
	} else if a <= 10 {
		return 1
	} else {
		return 0
	}
}

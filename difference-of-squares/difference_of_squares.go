package diffsquares

import "math"

func SquareOfSum(n int) int {
	res := 0

	for i := 1; i <= n; i++ {
		res += i
	}
	return int(math.Pow(float64(res), 2))
}

func SumOfSquares(n int) int {
	res := 0

	for i := 1; i <= n; i++ {
		res += int(math.Pow(float64(i), 2))
	}
	return res
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

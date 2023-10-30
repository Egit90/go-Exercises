package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {

	stringNumber := strconv.Itoa(n)
	numberOfDigits := len(stringNumber)

	res := 0

	for _, v := range stringNumber {
		res += int(math.Pow(float64(v-'0'), float64(numberOfDigits)))
	}
	return res == n
}

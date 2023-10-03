package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	stepCount := 0
	if n <= 0 {
		return 0, errors.New("Error")
	}
	for n != 1 {
		stepCount++
		switch n % 2 {
		case 0:
			n = n / 2
		default:
			n = 3*n + 1
		}
	}
	return stepCount, nil
}

package prime

import "errors"

var ErrInvalidNthPrime = errors.New("n has to be positive")

func IsPrime(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, ErrInvalidNthPrime
	}
	primeCount := 0
	for i := 1; i < 1000000; i++ {
		if IsPrime(i) {
			primeCount++
		}
		if primeCount == n {
			return i, nil
		}
	}
	return 0, errors.New("not supported range")
}

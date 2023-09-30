package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, number int) (float64, error) {
	a, err1 := f.FodderAmount(number)
	if err1 != nil {
		return 0, err1
	}
	b, err2 := f.FatteningFactor()
	if err2 != nil {
		return 0, err2
	}

	return a * b / float64(number), nil

}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, number int) (float64, error) {
	if number > 0 {
		return DivideFood(f, number)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf(
		"%d cows are invalid: %s", e.cows, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numberOFCows int) error {
	if numberOFCows < 0 {
		return &InvalidCowsError{
			cows:    numberOFCows,
			message: "there are no negative cows",
		}
	} else if numberOFCows == 0 {
		return &InvalidCowsError{
			cows:    0,
			message: "no cows don't need food",
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

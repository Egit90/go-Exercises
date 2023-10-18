package lsproduct

import (
	"errors"
	"regexp"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var largestProduct int64

	re := regexp.MustCompile("^[0-9]+$")

	if !re.MatchString(digits) {
		return 0, errors.New("digits input must only contain digits")
	}

	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}

	if span <= 0 {
		return 0, errors.New("span must not be negative")
	}

	for i := 0; i < len(digits)-1; i++ {
		tmp := 1

		if span+i > len(digits) {
			break
		}

		current := digits[i : span+i]
		for _, myRune := range current {
			tmp *= int(myRune - '0')
		}

		if int64(tmp) > largestProduct {
			largestProduct = int64(tmp)
		}
	}
	return largestProduct, nil
}

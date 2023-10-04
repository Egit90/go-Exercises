package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if a == b {
		return 0, nil
	} else if len(a) != len(b) {
		return 0, errors.New("strings are not the same size")
	} else {
		res := 0
		for i := range a {
			if b[i] != a[i] {
				res++
			}
		}
		return res, nil
	}
}

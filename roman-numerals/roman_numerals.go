package romannumerals

import (
	"errors"
	"fmt"
)

var romanMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		er := fmt.Sprintf("%d is out of range", input)
		return "0", errors.New(er)
	}
	list := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := ""
	for i := 0; i < len(list); i++ {
		v := list[i]
		if v <= input {
			res += romanMap[v]
			input -= v
			i--
		}

	}
	return res, nil
}

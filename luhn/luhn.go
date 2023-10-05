package luhn

import (
	"regexp"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 || checkIfInputIncludesNonDigit(id) {
		return false
	}

	return luhen(id)
}

func luhen(input string) bool {
	result := 0
	modify := false
	for i := len(input) - 1; i >= 0; i-- {
		value := int(input[i] - '0')

		if modify {
			value = value * 2
			if value > 9 {
				value -= 9
			}
			modify = false
		} else {
			modify = true
		}
		result += int(value)
	}
	return result%10 == 0
}

func checkIfInputIncludesNonDigit(input string) bool {

	pattern := "[^0-9]"
	matched, _ := regexp.MatchString(pattern, input)
	return matched
}

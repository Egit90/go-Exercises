package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	isQustion := isItQuestion(remark)
	isYelling := isItYell(remark)
	if isQustion && !isYelling {
		return "Sure."
	} else if isQustion && isYelling {
		return "Calm down, I know what I'm doing!"
	} else if isYelling {
		return "Whoa, chill out!"
	} else if isItFineByHim(remark) {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}

func isItQuestion(input string) bool {
	input = strings.TrimSpace(input)
	return len(input) > 0 && input[len(input)-1] == '?'
}

func isItYell(input string) bool {
	haveLetter := strings.IndexFunc(input, unicode.IsLetter) >= 0
	capital := strings.ToUpper(input) == input
	return haveLetter && capital

}

func isItFineByHim(input string) bool {
	return strings.TrimSpace(input) == ""
}

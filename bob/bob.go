package bob

import (
	"strings"
	"unicode"
)

type Remark struct {
	remark string
}

func (r Remark) isItQuestion() bool {
	return len(r.remark) > 0 && r.remark[len(r.remark)-1] == '?'
}
func (r Remark) isItYell() bool {
	haveLetter := strings.IndexFunc(r.remark, unicode.IsLetter) >= 0
	capital := strings.ToUpper(r.remark) == r.remark
	return haveLetter && capital

}
func (r Remark) isItFineByHim() bool {
	return strings.TrimSpace(r.remark) == ""
}

func Hey(remark string) string {
	r := Remark{strings.TrimSpace(remark)}
	isQuestion := r.isItQuestion()
	isYelling := r.isItYell()

	if isQuestion && !isYelling {
		return "Sure."
	} else if isQuestion && isYelling {
		return "Calm down, I know what I'm doing!"
	} else if isYelling {
		return "Whoa, chill out!"
	} else if r.isItFineByHim() {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}

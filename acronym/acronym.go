package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	var sb strings.Builder
	put := true
	for _, v := range strings.ToUpper(s) {
		if v == ' ' || v == '-' {
			put = true
		}
		if put && unicode.IsLetter(v) {
			sb.WriteRune(v)
			put = false
		}

	}
	return sb.String()

}

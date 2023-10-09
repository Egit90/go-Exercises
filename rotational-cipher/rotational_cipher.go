package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	res := ""
	for _, v := range plain {
		if unicode.IsLetter(v) {
			letter := v + rune(shiftKey)
			if unicode.IsLower(v) && letter > 122 || unicode.IsUpper(v) && letter > 90 {
				letter -= 26
			}
			res += string(letter)
		} else {
			res += string(v)
		}
	}
	return res
}

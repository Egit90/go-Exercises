package isbn

import "strings"

func IsValidISBN(isbn string) bool {
	res := 0
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	for i, v := range isbn {
		runeValue := 0
		if v == rune('X') && i == len(isbn)-1 {
			runeValue = 10
		} else {
			runeValue = int(v - '0')
		}
		k := runeValue * (len(isbn) - i)
		res += k
	}
	return res%11 == 0
}

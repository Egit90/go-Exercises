package atbash

import (
	"bytes"
	"strings"
	"unicode"
)

func createCiperMap() map[rune]rune {
	Plain := "abcdefghijklmnopqrstuvwxyz"
	Cipher := "zyxwvutsrqponmlkjihgfedcba"
	myMap := map[rune]rune{}

	for i, v := range Plain {
		myMap[v] = rune(Cipher[i])
	}
	return myMap
}

func Atbash(s string) string {
	count := 0
	cipherMap := createCiperMap()
	var b bytes.Buffer

	for _, r := range strings.ToLower(s) {
		v, ok := cipherMap[r]
		if ok {
			b.WriteRune(v)
			count++
		}
		if !ok && unicode.IsDigit(r) {
			b.WriteRune(r)
			count++
		}
		if count == 5 {
			b.WriteRune(' ')
			count = 0
		}
	}
	return strings.Trim(b.String(), " ")
}

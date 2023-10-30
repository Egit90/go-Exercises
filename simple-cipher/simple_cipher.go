package cipher

import (
	"errors"
	"strings"
	"unicode"
)

type shift struct {
	distance int
}

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return &shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if distance <= -26 || distance == 0 || distance >= 26 {
		return nil
	}
	return shift{distance: distance}
}

func encodeRune(r rune, dis int32) (rune, error) {
	if unicode.IsLetter(r) {
		key := dis + r

		if key < 'a' {
			key += 26
		}
		if key > 'z' {
			key -= 26
		}

		return key, nil

	}
	return 0, errors.New("not a letter")
}

func (c shift) Encode(input string) string {
	var sb strings.Builder

	for _, v := range strings.ToLower(input) {
		letter, err := encodeRune(v, int32(c.distance))
		if err == nil {
			sb.WriteRune(letter)
		}
	}
	return sb.String()
}

func (c shift) Decode(input string) string {
	c.distance *= -1
	return c.Encode(input)
}

func NewVigenere(key string) Cipher {

	if len(key) < 3 {
		return nil
	}
	for _, r := range key {
		if !unicode.IsLetter(r) || unicode.IsUpper(r) {
			return nil
		}
	}
	return vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	count := 0
	var sb strings.Builder
	for _, l := range strings.ToLower(input) {
		if count == len(v.key) {
			count = 0
		}
		letter, err := encodeRune(l, int32(v.key[count]-'a'))
		if err == nil {
			sb.WriteRune(letter)
			count++
		}
	}
	return sb.String()
}

func (v vigenere) Decode(input string) string {
	count := 0
	var sb strings.Builder
	for _, l := range strings.ToLower(input) {
		if count == len(v.key) {
			count = 0
		}
		letter, err := encodeRune(l, -1*int32(v.key[count]-'a'))
		if err == nil {
			sb.WriteRune(letter)
			count++
		}
	}
	return sb.String()
}

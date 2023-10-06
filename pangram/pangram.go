package pangram

import "strings"

func createAlphabetsMap() map[string]int {
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	alphabetMap := make(map[string]int)

	for _, letter := range alphabets {
		alphabetMap[string(letter)] = 0
	}

	return alphabetMap

}

func IsPangram(input string) bool {
	myAlphabetMap := createAlphabetsMap()

	for _, v := range input {
		myAlphabetMap[strings.ToLower(string(v))]++
	}
	for _, v := range myAlphabetMap {
		if v == 0 {
			return false
		}
	}
	return true
}

package isogram

import "unicode"

func IsIsogram(word string) bool {
	myMap := map[rune]int{}
	for _, v := range word {
		if isItInMap(myMap, v) {
			return false
		}
		myMap[unicode.ToUpper(v)] = 1
	}
	return true
}

func isItInMap(myMap map[rune]int, MyRune rune) bool {
	_, ok := myMap[unicode.ToUpper(MyRune)]
	return ok && MyRune != rune('-') && MyRune != rune(' ')
}

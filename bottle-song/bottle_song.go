package bottlesong

import (
	"fmt"
	"strings"
)

var numberToWordMap = map[int]string{
	0:  "no",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Tweleve",
}

func Recite(startBottles, takeDown int) []string {
	res := []string{}
	for takeDown > 0 {
		res = append(res, Verse(startBottles, false))
		res = append(res, Verse(startBottles, false))
		res = append(res, "And if one green bottle should accidentally fall,")
		startBottles--
		res = append(res, Verse(startBottles, true))
		res = append(res, "")

		takeDown--
	}
	res = res[:len(res)-1]
	return res
}

func Verse(number int, isLastVerse bool) string {

	myNumber := numberToWordMap[number]
	bottle := ""

	if number == 1 {
		bottle = "bottle"
	} else {
		bottle = "bottles"
	}

	if !isLastVerse {
		return fmt.Sprintf("%s green %s hanging on the wall,", myNumber, bottle)
	} else {
		return fmt.Sprintf("There'll be %s green %s hanging on the wall.", strings.ToLower(myNumber), bottle)
	}
}

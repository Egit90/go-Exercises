package resistorcolorduo

import (
	"fmt"
	"strconv"
)

var myMap = map[string]int{
	"black": 0, "brown": 1, "red": 2, "orange": 3, "yellow": 4, "green": 5, "blue": 6, "violet": 7, "grey": 8, "white": 9,
}

func Value(colors []string) int {
	res := ""
	for i := 0; i <= 1; i++ {
		res += fmt.Sprint(myMap[colors[i]])
	}
	inte, _ := strconv.Atoi(res)
	return inte
}

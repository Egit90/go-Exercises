package reverse

import "strings"

func Reverse(input string) string {
	inputList := strings.Split(input, "")
	for i, j := 0, len(inputList)-1; i < j; i, j = i+1, j-1 {
		inputList[i], inputList[j] = inputList[j], inputList[i]
	}
	return strings.Join(inputList, "")
}

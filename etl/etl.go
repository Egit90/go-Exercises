package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}

	for k, v := range in {
		for _, innerValue := range v {
			out[strings.ToLower(innerValue)] = k
		}
	}
	return out
}

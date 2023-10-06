package strain

func Keep[T any](list []T, filter func(T) bool) []T {
	var res []T
	for _, v := range list {
		if filter(v) {
			res = append(res, v)
		}
	}
	return res
}

func Discard[T any](list []T, filter func(T) bool) []T {
	var res []T
	for _, v := range list {
		if !filter(v) {
			res = append(res, v)
		}
	}
	return res
}

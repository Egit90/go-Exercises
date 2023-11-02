package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {

	for _, v := range s.Reverse() {
		initial = fn(v, initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	tmp := IntList{}
	for _, v := range s {
		if fn(v) {
			tmp = append(tmp, v)
		}
	}
	return tmp
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	tmp := IntList{}
	for _, v := range s {
		tmp = append(tmp, fn(v))
	}
	return tmp
}

func (s IntList) Reverse() IntList {

	n := s.Length()

	tmp := make(IntList, n)
	for i, v := range s {
		tmp[n-i-1] = v
	}
	return tmp
}

func (s IntList) Append(lst IntList) IntList {
	for _, item := range lst {
		s = append(s, item)
	}
	return s
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, v := range lists {
		s = s.Append(v)
	}
	return s
}

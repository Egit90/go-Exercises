package grains

import "errors"

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("invaild entree")
	}
	if number == 1 {
		return 1, nil
	}

	v, _ := Square(number - 1)
	return v * 2, nil
}

func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		g, _ := Square(i)
		total += g
	}
	return total
}

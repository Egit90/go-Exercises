package triangle

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	} else {
		return Sca
	}
}
func isTriangle(a, b, c float64) bool {
	if a == 0 || b == 0 || c == 0 {
		return false
	}
	res := a+b >= c &&
		b+c >= a &&
		a+c >= b

	return res
}

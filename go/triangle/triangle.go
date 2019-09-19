// Package triangle implements the KindFromSides method
package triangle

import (
	"math"
)

// Kind type for return of KindFromSides
type Kind string

// consts for different types of triangles
const (
	NaT Kind = "not a triangle"
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	var k Kind
	k = Sca

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	
	switch {
	case math.IsInf(c, 1): //Used for case of side being +Inf
		k = NaT
	case a <= 0:
		k = NaT
	case a + b < c: // triangle inequality
		k = NaT
	case a == c:
		k = Equ
	case a == b || b == c:
		k = Iso
	}
	return k
}

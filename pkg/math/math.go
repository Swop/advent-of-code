package math

import "math"

// GCD computes the greatest common divisor (Euclidean algorithm)
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM computes the least common multiplier (via GCD method)
func LCM(n ...int) int {
	l := len(n)
	switch l {
	case 0:
		return 0
	case 1:
		return n[0]
	}
	lcm := func(a, b int) int {
		return a * b / GCD(a, b)
	}
	result := lcm(n[l-2], n[l-1])
	for i := 0; i < l-2; i++ {
		result = lcm(n[l-3-i], result)
	}
	return result
}

func ManhattanDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

// Shoelace computes the area of a polygon given its vertices
// https://en.wikipedia.org/wiki/Shoelace_formula
//
// Can be coupled with the Pick's theorem
// https://en.wikipedia.org/wiki/Pick%27s_theorem
// A = i + b/2 - 1
// where `A` is the area of the polygon,
// `i` is the number of lattice points in the interior of the polygon,
// and `b` is the number of lattice points on the boundary of the polygon.
func Shoelace(vertices [][2]int) float64 {
	total := 0
	for i := 0; i < len(vertices); i++ {
		v1 := vertices[i]
		v2 := vertices[(i+1)%len(vertices)]

		total += (v1[0] * v2[1]) - (v1[1] * v2[0])
	}
	return math.Abs(float64(total)) / 2
}

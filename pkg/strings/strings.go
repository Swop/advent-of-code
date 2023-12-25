package strings

import (
	"strconv"
	"strings"
)

func ToInts(s string, sep string) []int {
	var ints []int
	for _, val := range strings.Split(s, sep) {
		ints = append(ints, ToInt(val))
	}
	return ints
}

func To2Int(s string, sep string) [2]int {
	var ints [2]int
	for i, val := range strings.Split(s, sep) {
		ints[i] = ToInt(val)
	}
	return ints
}

func To3Int(s string, sep string) [3]int {
	var ints [3]int
	for i, val := range strings.Split(s, sep) {
		ints[i] = ToInt(val)
	}
	return ints
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ToFloat64(s string) float64 {
	return float64(ToInt(s))
}

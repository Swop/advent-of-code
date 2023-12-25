package slices

import (
	"cmp"
	"strconv"
)

func Intersect[T comparable](slice1, slice2 []T) []T {
	var intersect []T
	for _, element1 := range slice1 {
		for _, element2 := range slice2 {
			if element1 == element2 {
				intersect = append(intersect, element1)
			}
		}
	}
	return intersect
}

func Reverse[T any](slice []T) []T {
	reversed := make([]T, len(slice))
	for i, element := range slice {
		reversed[len(slice)-1-i] = element
	}
	return reversed
}

func IntsToStrings(ints []int) []string {
	strings := make([]string, 0, len(ints))
	for _, i := range ints {
		strings = append(strings, strconv.Itoa(i))
	}
	return strings
}

func Zip[T any](slices ...[]T) [][]T {
	var zipped [][]T
	lengths := make([]int, len(slices))
	for i, slice := range slices {
		lengths[i] = len(slice)
	}
	m := SliceMin(lengths)
	for i := 0; i < m; i++ {
		var z []T
		for _, slice := range slices {
			z = append(z, slice[i])
		}
		zipped = append(zipped, z)
	}
	return zipped
}

func Combinations[T any](iterable []T, r int) [][]T {
	pool := iterable
	n := len(pool)

	if r > n {
		return nil
	}

	indices := make([]int, r)
	for i := range indices {
		indices[i] = i
	}

	result := make([]T, r)
	for i, el := range indices {
		result[i] = pool[el]
	}
	s2 := make([]T, r)
	copy(s2, result)
	var rt [][]T
	rt = append(rt, s2)

	for {
		i := r - 1
		for {
			if i < 0 || indices[i] != i+n-r {
				break
			}
			i--
		}

		if i < 0 {
			return rt
		}

		indices[i]++
		for j := i + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}

		for ; i < len(indices); i++ {
			result[i] = pool[indices[i]]
		}
		s2 = make([]T, r)
		copy(s2, result)
		rt = append(rt, s2)
	}
}

func SliceMin[T cmp.Ordered](slice []T) T {
	m := slice[0]
	for _, element := range slice {
		if element < m {
			m = element
		}
	}
	return m
}

func SliceMax[T cmp.Ordered](slice []T) T {
	m := slice[0]
	for _, element := range slice {
		if element > m {
			m = element
		}
	}
	return m
}

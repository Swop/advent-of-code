package set

import (
	"fmt"

	"github.com/Swop/advent-of-code/pkg/maps"
)

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return Set[T]{}
}

func NewWithValues[T comparable](values ...T) Set[T] {
	s := New[T]()
	for _, v := range values {
		s.Add(v)
	}
	return s
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Size() int {
	return len(maps.MapKeys(s))
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) ContainsAll(slice []T) bool {
	for _, v := range slice {
		if !s.Has(v) {
			return false
		}
	}
	return true
}

func (s Set[T]) ContainsAny(slice []T) bool {
	for _, v := range slice {
		if s.Has(v) {
			return true
		}
	}
	return false
}

func (s Set[T]) IsSubset(other Set[T]) bool {
	return other.ContainsAll(maps.MapKeys(s))
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Values() []T {
	return maps.MapKeys(s)
}

func (s Set[T]) String() string {
	return fmt.Sprintf("%v", s.Values())
}

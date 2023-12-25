package test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/Swop/advent-of-code/pkg/runner"
	"github.com/google/go-cmp/cmp"
)

type Table map[string]Test

type Test struct {
	Func  runner.PartFunc
	Input string
	Want  any
}

func splitInput(input string) []string {
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func PuzzleTest(t *testing.T, tt Table) {
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			diff := cmp.Diff(tc.Func(splitInput(tc.Input)), tc.Want)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func PuzzleBenchmark(b *testing.B, input string, parts ...runner.PartFunc) {
	lines := splitInput(input)
	for i, f := range parts {
		b.Run("part_"+strconv.Itoa(i+1), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				f(lines)
			}
		})
	}
}

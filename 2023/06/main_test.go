package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay6(t *testing.T) {
	sample := `Time:      7  15   30
Distance:  9  40  200`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  288,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  71503,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  588588,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  34655848,
		},
	})
}

func BenchmarkDay6(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

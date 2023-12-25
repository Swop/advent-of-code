package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay9(t *testing.T) {
	sample := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  114,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  2,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  1992273652,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  1012,
		},
	})
}

func BenchmarkDay9(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

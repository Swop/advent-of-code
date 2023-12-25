package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDayX(t *testing.T) {
	sample := ``
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  0,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  0,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  0,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  0,
		},
	})
}

func BenchmarkDayX(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

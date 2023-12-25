package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay24(t *testing.T) {
	sample := `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`
	p1 := func(area [2]float64) func(input []string) any {
		return func(input []string) any { return resolvePart1(input, area) }
	}
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  p1([2]float64{7, 27}),
			Input: sample,
			Want:  2,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  47,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  17906,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  571093786416929,
		},
	})
}

func BenchmarkDay24(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay13(t *testing.T) {
	sample := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  405,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  400,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  37113,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  30449,
		},
	})
}

func BenchmarkDay13(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

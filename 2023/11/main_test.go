package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay11(t *testing.T) {
	sample := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  374,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  82000210,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  10292708,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  790194712336,
		},
	})
}

func BenchmarkDay11(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

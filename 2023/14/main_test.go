package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay14(t *testing.T) {
	sample := `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  136,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  64,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  113486,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  104409,
		},
	})
}

func BenchmarkDay14(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

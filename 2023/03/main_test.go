package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay3(t *testing.T) {
	sample := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	test.PuzzleTest(t, test.Table{
		"part 1 sample": {
			Func:  part1,
			Input: sample,
			Want:  4361,
		},
		"part 2 sample": {
			Func:  part2,
			Input: sample,
			Want:  467835,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  537732,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  84883664,
		},
	})
}

func BenchmarkDay3(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

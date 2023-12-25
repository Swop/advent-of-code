package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay18(t *testing.T) {
	sample := `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  62,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  952408144115,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  70253,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  131265059885080,
		},
	})
}

func BenchmarkDay18(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

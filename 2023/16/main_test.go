package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay16(t *testing.T) {
	sample := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  46,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  51,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  7477,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  7853,
		},
	})
}

func BenchmarkDay16(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

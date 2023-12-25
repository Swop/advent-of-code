package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay7(t *testing.T) {
	sample := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  6440,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  5905,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  246163188,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  245794069,
		},
	})
}

func BenchmarkDay7(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

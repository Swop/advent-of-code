package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay4(t *testing.T) {
	sample := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  13,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  30,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  23441,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  5923918,
		},
	})
}

func BenchmarkDay4(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

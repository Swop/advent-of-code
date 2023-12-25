package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay2(t *testing.T) {
	sample := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  8,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  2286,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  2085,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  79315,
		},
	})
}

func BenchmarkDay2(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

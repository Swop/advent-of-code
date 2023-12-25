package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay1(t *testing.T) {
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func: part1,
			Input: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			Want: 142,
		},
		"part_2_sample": {
			Func: part2,
			Input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			Want: 281,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  54916,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  54728,
		},
	})
}

func BenchmarkDay1(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

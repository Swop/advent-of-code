package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay20(t *testing.T) {
	test.PuzzleTest(t, test.Table{
		"part_1_sample_1": {
			Func: part1,
			Input: `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`,
			Want: 32000000,
		},
		"part_1_sample_2": {
			Func: part1,
			Input: `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`,
			Want: 11687500,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  818649769,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  246313604784977,
		},
	})
}

func BenchmarkDay20(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

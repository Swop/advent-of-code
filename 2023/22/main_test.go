package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay22(t *testing.T) {
	sample := `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  5,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  7,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  519,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  109531,
		},
	})
}

func BenchmarkDay22(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay15(t *testing.T) {
	sample := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  1320,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  145,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  517015,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  286104,
		},
	})
}

func BenchmarkDay15(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

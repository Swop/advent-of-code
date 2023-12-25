package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay17(t *testing.T) {
	sample := `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  102,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  94,
		},
		"part_2_sample_2": {
			Func: part2,
			Input: `111111111111
999999999991
999999999991
999999999991
999999999991`,
			Want: 71,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  907,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  1057,
		},
	})
}

func BenchmarkDay17(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

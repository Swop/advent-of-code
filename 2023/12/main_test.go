package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay12(t *testing.T) {
	sample := `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  21,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  525152,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  7379,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  7732028747925,
		},
	})
}

func BenchmarkDay12(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

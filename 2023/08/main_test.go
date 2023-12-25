package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestPuzzle(t *testing.T) {
	test.PuzzleTest(t, test.Table{
		"part_1_sample_1": {
			Func: part1,
			Input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
			Want: 2,
		},
		"part_1_sample_2": {
			Func: part1,
			Input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`,
			Want: 6,
		},
		"part_2_sample": {
			Func: part2,
			Input: `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`,
			Want: 6,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  12169,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  12030780859469,
		},
	})
}

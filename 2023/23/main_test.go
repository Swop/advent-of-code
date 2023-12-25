package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay23(t *testing.T) {
	sample := `#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  94,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  154,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  1966,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  6286,
		},
	})
}

func BenchmarkDay23(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

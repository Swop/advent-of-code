package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay21(t *testing.T) {
	sample := `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
`
	p1 := func(maxSteps int) func(input []string) any {
		return func(input []string) any { return resolvePart1(input, maxSteps) }
	}
	// ---------------------
	// See comment below
	// ---------------------
	// p2 := func(maxSteps int) func(input []string) any {
	//	return func(input []string) any { return resolvePart2(input, maxSteps) }
	//}
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  p1(6),
			Input: sample,
			Want:  16,
		},
		// ---------------------
		// Since the part2 solver has been developed using unique characteristic of the input (centered start, perfect square,
		// blank horizontal & vertical lines at level of start point, ...) which are not present in the sample,
		// we can't assess the correctness of the solver on the sample input with the different
		// [max step => expected results] given in the challenge page. Therefore, the following tests has been commented out.
		// ---------------------
		// "part_2_sample_1": {
		//	Func:  p2(6),
		//	Input: sample,
		//	Want:  16,
		// },
		// "part_2_sample_2": {
		//	Func:  p2(10),
		//	Input: sample,
		//	Want:  50,
		// },
		// "part_2_sample_3": {
		//	Func:  p2(50),
		//	Input: sample,
		//	Want:  1594,
		// },
		// "part_2_sample_4": {
		//	Func:  p2(100),
		//	Input: sample,
		//	Want:  6536,
		// },
		// "part_2_sample_5": {
		//	Func:  p2(500),
		//	Input: sample,
		//	Want:  167004,
		// },
		// "part_2_sample_6": {
		//	Func:  p2(1_000),
		//	Input: sample,
		//	Want:  668697,
		// },
		// "part_2_sample_7": {
		//	Func:  p2(5_000),
		//	Input: sample,
		//	Want:  16733044,
		// },
		// "part_1_input": {
		//	Func:  part1,
		//	Input: input,
		//	Want:  3768,
		// },
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  627960775905777,
		},
	})
}

func BenchmarkDay21(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

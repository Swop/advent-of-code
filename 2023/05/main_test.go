package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay5(t *testing.T) {
	sample := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  35,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  46,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  199602917,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  2254686,
		},
	})
}

func BenchmarkDay5(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

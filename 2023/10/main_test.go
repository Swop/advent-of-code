package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay10(t *testing.T) {
	test.PuzzleTest(t, test.Table{
		"part_1_sample_1": {
			Func: part1,
			Input: `.....
.S-7.
.|.|.
.L-J.
.....`,
			Want: 4,
		},
		"part_1_sample_2": {
			Func: part1,
			Input: `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`,
			Want: 8,
		},
		"part_2_sample_1": {
			Func: part2,
			Input: `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`,
			Want: 4,
		},
		"part_2_sample_2": {
			Func: part2,
			Input: `..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........`,
			Want: 4,
		},
		"part_2_sample_3": {
			Func: part2,
			Input: `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`,
			Want: 8,
		},
		"part_2_sample_4": {
			Func: part2,
			Input: `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`,
			Want: 10,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  6909,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  461,
		},
	})
}

func BenchmarkDay10(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

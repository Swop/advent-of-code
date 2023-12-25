package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay19(t *testing.T) {
	sample := `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  part1,
			Input: sample,
			Want:  19114,
		},
		"part_2_sample": {
			Func:  part2,
			Input: sample,
			Want:  167409079868000,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  362930,
		},
		"part_2_input": {
			Func:  part2,
			Input: input,
			Want:  116365820987729,
		},
	})
}

func BenchmarkDay19(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1, part2)
}

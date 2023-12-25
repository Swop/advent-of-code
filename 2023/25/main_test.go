package main

import (
	_ "embed"
	"testing"

	"github.com/Swop/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

func TestDay25(t *testing.T) {
	sample := `jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr`
	p1 := func(edgeToRemove [3][2]string) func(input []string) any {
		return func(input []string) any { return resolvePart1(input, edgeToRemove) }
	}
	test.PuzzleTest(t, test.Table{
		"part_1_sample": {
			Func:  p1([3][2]string{{"cmg", "bvb"}, {"jqt", "nvd"}, {"pzl", "hfx"}}),
			Input: sample,
			Want:  54,
		},
		"part_1_input": {
			Func:  part1,
			Input: input,
			Want:  507626,
		},
	})
}

func BenchmarkDay25(b *testing.B) {
	test.PuzzleBenchmark(b, input, part1)
}

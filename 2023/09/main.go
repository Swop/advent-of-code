package main

import (
	"strconv"
	"strings"

	"github.com/Swop/advent-of-code/pkg/runner"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	return compute(parseInput(1, input))
}

func part2(input []string) any {
	return compute(parseInput(2, input))
}

func compute(histories [][]int) any {
	total := 0
	for _, h := range histories {
		total += seq(h)
	}
	return total
}

func seq(h []int) int {
	s := make([]int, len(h)-1)
	sameGap := true
	var gap int

	for i := 0; i < len(h)-1; i++ {
		g := h[i+1] - h[i]
		if g != gap {
			sameGap = false
		}
		gap = g
		s[i] = gap
	}

	if sameGap {
		return h[len(h)-1] + gap
	}
	return h[len(h)-1] + seq(s)
}

func parseInput(part int, input []string) [][]int {
	var histories [][]int
	for _, line := range input {
		var h []int
		f := strings.Fields(line)
		if part == 1 {
			for _, s := range f {
				n, _ := strconv.Atoi(s)
				h = append(h, n)
			}
		} else {
			for i := len(f) - 1; i >= 0; i-- {
				n, _ := strconv.Atoi(f[i])
				h = append(h, n)
			}
		}
		histories = append(histories, h)
	}
	return histories
}

package main

import (
	"github.com/Swop/advent-of-code/pkg/runner"
	"github.com/Swop/advent-of-code/pkg/slices"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	return compute(parseInput(input), 0)
}

func part2(input []string) any {
	return compute(parseInput(input), 1)
}

func compute(patterns [][]string, dist int) any {
	total := 0
	for _, p := range patterns {
		total += value(p, dist)*100 + value(rotate(p), dist)
	}
	return total
}

func value(p []string, dist int) int {
	v := 0
	for i := 0; i < len(p)-1; i++ {
		totalDist := 0
		z := slices.Zip(p[i+1:], slices.Reverse(p[:i+1]))
		for _, zipped := range z {
			totalDist += distance(zipped[0], zipped[1])
		}
		if totalDist == dist {
			v += i + 1
		}
	}
	return v
}

func distance(s1, s2 string) int {
	zipped := slices.Zip([]uint8(s1), []uint8(s2))
	sum := 0
	for _, z := range zipped {
		if z[0] != z[1] {
			sum++
		}
	}
	return sum
}

func rotate(p []string) []string {
	rotated := make([]string, len(p[0]))
	for i := 0; i < len(p[0]); i++ {
		var s string
		for _, row := range p {
			s += string(row[i])
		}
		rotated[i] = s
	}
	return rotated
}

func parseInput(input []string) [][]string {
	patterns := make([][]string, 0)
	currentPattern := make([]string, 0)
	for _, line := range input {
		if line == "" {
			patterns = append(patterns, currentPattern)
			currentPattern = make([]string, 0)
			continue
		}
		currentPattern = append(currentPattern, line)
	}
	patterns = append(patterns, currentPattern)
	return patterns
}

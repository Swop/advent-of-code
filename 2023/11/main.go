package main

import (
	"github.com/Swop/advent-of-code/pkg/math"
	"github.com/Swop/advent-of-code/pkg/runner"
	"github.com/Swop/advent-of-code/pkg/slices"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	return compute(parseInput(1, input))
}

func part2(input []string) any {
	return compute(parseInput(999_999, input))
}

func compute(galaxies []*[2]int) any {
	comb := slices.Combinations(galaxies, 2)
	total := 0
	for _, galCouple := range comb {
		total += math.ManhattanDistance(galCouple[0][0], galCouple[0][1], galCouple[1][0], galCouple[1][1])
	}
	return total
}

func parseInput(expansion int, input []string) []*[2]int {
	h := len(input)
	w := len(input[0])
	xIdx := make([]int, w)
	yIdx := make([]int, h)
	for x := 0; x < w; x++ {
		xIdx[x] = 1 + expansion
	}
	for y := 0; y < h; y++ {
		yIdx[y] = 1 + expansion
	}
	galaxies := make([]*[2]int, 0, 100)
	name := 1
	for y, line := range input {
		for x, c := range line {
			if c == '.' {
				continue
			}
			galaxies = append(galaxies, &[2]int{x, y})
			xIdx[x] = 1
			yIdx[y] = 1
			name++
		}
	}

	for x := 1; x < w; x++ {
		xIdx[x] += xIdx[x-1]
	}
	for y := 1; y < h; y++ {
		yIdx[y] += yIdx[y-1]
	}

	for _, gal := range galaxies {
		gal[0] = xIdx[gal[0]]
		gal[1] = yIdx[gal[1]]
	}

	return galaxies
}

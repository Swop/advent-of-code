package main

import (
	"github.com/Swop/advent-of-code/pkg/grid"
	"github.com/Swop/advent-of-code/pkg/runner"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	g := parseInput(input)
	tilt(g)
	return computeLoad(g)
}

func part2(input []string) any {
	g := parseInput(input)
	cache := make(map[string]int)
	cycleCount := 0
	var total []int
	for {
		cycleCount++
		cycle(g)
		total = append(total, computeLoad(g))
		h := g.Hash()
		if v, ok := cache[h]; ok {
			diff := cycleCount - v
			c := total[len(total)-diff:]
			return c[((1_000_000_000-cycleCount)%diff)-1]
		}
		cache[h] = cycleCount
	}
}

func tilt(g *grid.Grid[rune]) {
	for x := 0; x < g.Width(); x++ {
		for y := g.Height() - 1; y >= 0; y-- {
			c := g.Get(x, y)
			if c == 'O' {
				lastValidPlace := y
				i := y - 1
				if i < 0 {
					continue
				}
				ci := g.Get(x, i)
				for i >= 0 && ci != '#' {
					if ci == '.' {
						lastValidPlace = i
					}
					i--
					if i >= 0 {
						ci = g.Get(x, i)
					}
				}
				if y != lastValidPlace {
					g.Set(x, y, '.')
					g.Set(x, lastValidPlace, 'O')
				}
			}
		}
	}
}

func cycle(g *grid.Grid[rune]) {
	for i := 0; i < 4; i++ {
		tilt(g)
		g.Rotate(-1)
	}
}

func computeLoad(g *grid.Grid[rune]) int {
	load := 0
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if g.Get(x, y) == 'O' {
				load += g.Height() - y
			}
		}
	}
	return load
}

func parseInput(input []string) *grid.Grid[rune] {
	g := grid.New[rune](len(input[0]), len(input), func(r rune, _ [2]int) string {
		return string(r)
	})
	for y, line := range input {
		for x, c := range line {
			g.Set(x, y, c)
		}
	}
	return g
}

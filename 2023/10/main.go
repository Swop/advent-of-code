package main

import (
	"github.com/Swop/advent-of-code/pkg/grid"
	"github.com/Swop/advent-of-code/pkg/math"
	"github.com/Swop/advent-of-code/pkg/runner"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	_, perimeter := getVerticesAndPerimeter(parseInput(input))
	return perimeter / 2
}

func part2(input []string) any {
	vertices, perimeter := getVerticesAndPerimeter(parseInput(input))
	// Shoelace formula + Pick's theorem
	return int(math.Shoelace(vertices)) - perimeter/2 + 1
}

func getVerticesAndPerimeter(g *grid.Grid[rune], start [2]int) ([][2]int, int) {
	var vertices [][2]int
	perimeter := 1
	n := validNeighbours(g, start, start)
	if n[0].Vector[0] != n[1].Vector[0] || n[0].Vector[1] != n[1].Vector[1] {
		// start is a corner
		vertices = append(vertices, start)
	}
	previous := start
	current := n[0].Position
	isCorner := func(r rune) bool { return r == '7' || r == 'L' || r == 'J' || r == 'F' }
	for {
		if current == start {
			break
		}
		perimeter++
		if isCorner(g.Get(current[0], current[1])) {
			vertices = append(vertices, current)
		}
		n := validNeighbours(g, current, previous)
		previous = current
		current = n[0].Position
	}
	return vertices, perimeter
}

func validNeighbours(g *grid.Grid[rune], p [2]int, exclude [2]int) []grid.Neighbor {
	valid := make([]grid.Neighbor, 0, 2)
	pV := g.Get(p[0], p[1])
	for _, n := range g.Neighbors(p, false) {
		nV := g.Get(n.Position[0], n.Position[1])
		if nV == '.' || n.Position == exclude {
			continue
		}
		if (n.Vector[1] == -1 && (nV == '7' || nV == '|' || nV == 'F' || nV == 'S') && (pV == 'L' || pV == '|' || pV == 'J' || pV == 'S')) ||
			(n.Vector[0] == -1 && (nV == 'L' || nV == '-' || nV == 'F' || nV == 'S') && (pV == 'J' || pV == '-' || pV == '7' || pV == 'S')) ||
			(n.Vector[0] == 1 && (nV == 'J' || nV == '-' || nV == '7' || nV == 'S') && (pV == 'L' || pV == '-' || pV == 'F' || pV == 'S')) ||
			(n.Vector[1] == 1 && (nV == 'J' || nV == '|' || nV == 'L' || nV == 'S') && (pV == '7' || pV == '|' || pV == 'F' || pV == 'S')) {
			valid = append(valid, n)
		}
	}
	return valid
}

func parseInput(input []string) (*grid.Grid[rune], [2]int) {
	g := grid.New[rune](len(input[0]), len(input), func(r rune, _ [2]int) string {
		return string(r)
	})
	var start [2]int
	for y, line := range input {
		for x, c := range line {
			if c == 'S' {
				start = [2]int{x, y}
			}
			g.Set(x, y, c)
		}
	}
	return g, start
}

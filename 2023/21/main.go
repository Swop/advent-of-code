package main

import (
	"github.com/Swop/advent-of-code/pkg/grid"
	"github.com/Swop/advent-of-code/pkg/runner"
	"github.com/gammazero/deque"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	return resolvePart1(input, 64)
}

func part2(input []string) any {
	return resolvePart2(input, 26501365)
}

func resolvePart1(input []string, maxSteps int) any {
	g, start := parseInput(input)
	distances := findAllDistancesInSquare(g, start)

	parity := maxSteps % 2
	count := 0
	for _, d := range distances {
		if d%2 == parity && d <= maxSteps {
			count++
		}
	}
	return count
}

func resolvePart2(input []string, maxSteps int) any {
	g, start := parseInput(input)
	distances := findAllDistancesInSquare(g, start)
	// start in conveniently placed in the middle of the input grid, and grid is a perfect square.
	gridSize := g.Width()
	// in addition, the start point is surrounded by a vertical & horizontal blank line, so we can assume the distance
	// from the start to any edge (top/bottom/left/right) is equal to the distance from the start to the edge of the
	// grid (no obstacles).
	startToEdgeDist := gridSize / 2
	// count the number of even/odd distances, and the number of even/odd distances that are greater than the distance
	// from the start to the edge (i.e. "corners", because when we'll cover this distance and expand the grid, those
	// points will be placed outside the initial grid size).
	var evenCorners, oddCorners, evenFull, oddFull int
	for _, d := range distances {
		switch d % 2 {
		case 0:
			evenFull++
			if d > startToEdgeDist {
				evenCorners++
			}
		case 1:
			oddFull++
			if d > startToEdgeDist {
				oddCorners++
			}
		}
	}
	// number of square repetitions (in addition of the initial square)
	// this represents the number of horizontal expansions needed to walk the maxSteps in one direction from the start
	// point. This is possible because the start point is conveniently surrounded by horizontal & vertical blank lines.
	reps := (maxSteps - startToEdgeDist) / gridSize
	// When we reach maxSteps, we will consider only the points which match same parity (odd or even)
	var full, otherParityFull, corners, otherParityCorners int
	switch maxSteps % 2 {
	case 0:
		full = evenFull
		otherParityFull = oddFull
		corners = evenCorners
		otherParityCorners = oddCorners
	case 1:
		full = oddFull
		otherParityFull = evenFull
		corners = oddCorners
		otherParityCorners = evenCorners
	}
	return ((reps+1)*(reps+1))*full + (reps*reps)*otherParityFull - (reps+1)*corners + reps*otherParityCorners
}

// findAllDistancesInSquare computes the distance from the start point to all the points in the square (shortest path).
func findAllDistancesInSquare(g *grid.Grid[rune], start [2]int) map[[2]int]int {
	visited := map[[2]int]int{}
	type state struct {
		Pos  [2]int
		Dist int
	}
	q := deque.Deque[state]{}
	q.PushBack(state{
		Pos:  start,
		Dist: 0,
	})
	for q.Len() > 0 {
		s := q.PopFront()
		if _, ok := visited[s.Pos]; ok {
			continue
		}
		visited[s.Pos] = s.Dist
		for _, n := range g.Neighbors(s.Pos, false) {
			if g.Get(n.Position[0], n.Position[1]) == '#' {
				continue
			}
			q.PushBack(state{
				Pos:  n.Position,
				Dist: s.Dist + 1,
			})
		}
	}
	return visited
}

func parseInput(input []string) (*grid.Grid[rune], [2]int) {
	g := grid.NewFromInput(input)
	var start [2]int
	for _, e := range g.Enumerate() {
		if e.Value == 'S' {
			start = e.Position
			break
		}
	}
	return g, start
}

package main

import (
	"strconv"

	"github.com/Swop/advent-of-code/pkg/grid"
	"github.com/Swop/advent-of-code/pkg/heap"
	"github.com/Swop/advent-of-code/pkg/runner"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	g := parseInput(input)
	return compute(g, 1, 3)
}

func part2(input []string) any {
	g := parseInput(input)
	return compute(g, 4, 10)
}

type PathItem struct {
	Pos           [2]int
	Vector        [2]int
	HeatLoss      int
	CumulativeDir int
	Previous      *PathItem
}

type seeKey struct {
	pos           [2]int
	vector        [2]int
	cumulativeDir int
}

func compute(g *grid.Grid[int], min, max int) any {
	h := heap.New(func(i, j *PathItem) bool {
		return i.HeatLoss < j.HeatLoss
	})
	seen := make(map[seeKey]*PathItem)
	h.Push(&PathItem{
		Pos:           [2]int{0, 0},
		Vector:        [2]int{0, 0},
		HeatLoss:      0,
		CumulativeDir: 1,
		Previous:      nil,
	})

	for h.Size() > 0 {
		last := h.Pop()
		_, ok := seen[seeKey{pos: last.Pos, vector: last.Vector, cumulativeDir: last.CumulativeDir}]
		if ok {
			continue
		}
		seen[seeKey{pos: last.Pos, vector: last.Vector, cumulativeDir: last.CumulativeDir}] = last
		if last.Pos == [2]int{g.Width() - 1, g.Height() - 1} && last.CumulativeDir >= min {
			return last.HeatLoss
		}

		for _, n := range g.Neighbors(last.Pos, false) {
			if n.Vector[0] == -last.Vector[0] && n.Vector[1] == -last.Vector[1] {
				// reverse direction
				continue
			}
			cumulativeDir := 1
			switch {
			case n.Vector[0] == last.Vector[0] && n.Vector[1] == last.Vector[1]:
				// same direction
				cumulativeDir += last.CumulativeDir
				if cumulativeDir > max {
					continue
				}
			case last.Previous != nil && last.CumulativeDir < min:
				continue
			}
			pi := &PathItem{
				Pos:           n.Position,
				Vector:        n.Vector,
				HeatLoss:      last.HeatLoss + g.Get(n.Position[0], n.Position[1]),
				CumulativeDir: cumulativeDir,
				Previous:      last,
			}
			h.Push(pi)
		}
	}
	return 0
}

func parseInput(input []string) *grid.Grid[int] {
	g := grid.New[int](len(input[0]), len(input), func(r int, _ [2]int) string {
		return strconv.Itoa(r)
	})
	for y, line := range input {
		for x, c := range line {
			v, _ := strconv.Atoi(string(c))
			g.Set(x, y, v)
		}
	}
	return g
}

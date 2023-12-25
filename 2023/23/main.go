package main

import (
	"math"
	"slices"

	"github.com/Swop/advent-of-code/pkg/runner"
	"github.com/Swop/advent-of-code/pkg/set"
	"github.com/gammazero/deque"
)

func main() {
	runner.Run(part1, part2)
}

func part1(input []string) any {
	return compute(parseInput(1, input), len(input[0]), len(input))
}

func part2(input []string) any {
	return compute(toWeightedGraph(parseInput(2, input)), len(input[0]), len(input))
}

type Edges map[[2]int]set.Set[[3]int]

func toWeightedGraph(edges Edges) Edges {
	// we're going to reduce the complexity of the problem by shortening the path in corridors
	// (and introduced the notions of length, i.e. weighted edges)
	for {
		hasBroke := false
		for pos, neighbours := range edges {
			if neighbours.Size() == 2 {
				nSlice := neighbours.Values()
				n1X, n1Y, n1Length := nSlice[0][0], nSlice[0][1], nSlice[0][2]
				n1Pos := [2]int{n1X, n1Y}
				n2X, n2Y, n2Length := nSlice[1][0], nSlice[1][1], nSlice[1][2]
				n2Pos := [2]int{n2X, n2Y}
				// disconnect n1 and n2 from current tile, and connect them together (with the sum of their lengths)
				edges[n1Pos].Remove([3]int{pos[0], pos[1], n1Length})
				edges[n2Pos].Remove([3]int{pos[0], pos[1], n2Length})
				edges[n1Pos].Add([3]int{n2X, n2Y, n1Length + n2Length})
				edges[n2Pos].Add([3]int{n1X, n1Y, n1Length + n2Length})
				delete(edges, pos)
				hasBroke = true
				break
			}
		}
		if !hasBroke {
			break
		}
	}
	return edges
}

func compute(edges Edges, w, h int) any {
	visited := set.New[[2]int]()
	var maxPathLength float64
	q := deque.New[[3]int]()
	q.PushBack([3]int{1, 1, 1})
	for q.Len() > 0 {
		s := q.PopBack() // DFS
		x, y, pathLength := s[0], s[1], s[2]
		if pathLength == -1 {
			visited.Remove([2]int{x, y})
			continue
		}
		if x == w-2 && y == h-2 {
			maxPathLength = math.Max(maxPathLength, float64(pathLength))
			continue
		}
		if visited.Has([2]int{x, y}) {
			continue
		}
		visited.Add([2]int{x, y})
		q.PushBack([3]int{x, y, -1})
		// Perf improvement: check if neighbors contains the exit point. If so, we can stop the search
		// because other searches won't be able to join the exit because already visited
		var toPush [][3]int
		for _, next := range edges[[2]int{x, y}].Values() {
			if next[0] == w-2 && next[1] == h-2 {
				toPush = [][3]int{next}
				break
			}
			toPush = append(toPush, next)
		}

		for _, next := range toPush {
			q.PushBack([3]int{next[0], next[1], pathLength + next[2]})
		}
	}
	return int(maxPathLength) + 1 // +1 because we don't count the last step
}

func parseInput(part int, input []string) Edges {
	// patching the input to fix holes (avoids checking for out of bounds)
	input[0] = "##" + input[0][2:]
	input[len(input)-1] = input[len(input)-1][:len(input[len(input)-1])-2] + "##"

	edges := make(map[[2]int]set.Set[[3]int])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	addToSet := func(x, y, xx, yy, length int) {
		if _, ok := edges[[2]int{x, y}]; !ok {
			edges[[2]int{x, y}] = set.New[[3]int]()
		}
		edges[[2]int{x, y}].Add([3]int{xx, yy, length})
	}

	for y, line := range input {
		for x, char := range line {
			regularTileCharSet := []rune{'.'}
			if part == 2 {
				regularTileCharSet = []rune{'.', '>', 'v'}
			}
			switch {
			case slices.Contains(regularTileCharSet, char):
				for _, dir := range dirs {
					xx, yy := x+dir[0], y+dir[1]
					if slices.Contains(regularTileCharSet, rune(input[yy][xx])) {
						addToSet(x, y, xx, yy, 1)
						addToSet(xx, yy, x, y, 1)
					}
				}
			case part == 1 && char == '>':
				addToSet(x, y, x+1, y, 1)
				addToSet(x-1, y, x, y, 1)
			case part == 1 && char == 'v':
				addToSet(x, y, x, y+1, 1)
				addToSet(x, y-1, x, y, 1)
			}
			// note: both sample and input doesn't contain any ^ and <
		}
	}
	return edges
}

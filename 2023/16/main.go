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
	ch := make(chan int)
	go computeEnergized(grid.NewFromInput(input), [4]int{0, 0, 1, 0}, ch)
	return <-ch
}

func part2(input []string) any {
	g := grid.NewFromInput(input)
	ch := make(chan int)
	for y := 0; y < g.Height(); y++ {
		go computeEnergized(g, [4]int{0, y, 1, 0}, ch)
		go computeEnergized(g, [4]int{g.Width() - 1, y, -1, 0}, ch)
	}
	for x := 0; x < g.Width(); x++ {
		go computeEnergized(g, [4]int{x, 0, 0, 1}, ch)
		go computeEnergized(g, [4]int{x, g.Height() - 1, 0, -1}, ch)
	}

	maxEnergized := 0
	count := g.Width()*2 + g.Height()*2
	for i := 0; i < count; i++ {
		v := <-ch
		if v > maxEnergized {
			maxEnergized = v
		}
	}
	return maxEnergized
}

func computeEnergized(g *grid.Grid[rune], start [4]int, ch chan int) {
	q := deque.New[[4]int]()
	q.PushBack(start)
	seen := make(map[[4]int]struct{})
	energized := make(map[[2]int]struct{})
	for q.Len() > 0 {
		n := q.PopFront()
		seen[n] = struct{}{}
		energized[[2]int{n[0], n[1]}] = struct{}{}
		enqueueNext(g, q, seen, n)
	}
	ch <- len(energized)
}

func enqueueNext(g *grid.Grid[rune], q *deque.Deque[[4]int], seen map[[4]int]struct{}, n [4]int) {
	enq := func(x, y, dx, dy int) {
		next := [4]int{x, y, dx, dy}
		if _, ok := seen[next]; !ok && g.IsInGrid(next[0], next[1]) {
			q.PushBack(next)
		}
	}
	v := g.Get(n[0], n[1])
	if v == '.' || (v == '|' && n[2] == 0) || (v == '-' && n[3] == 0) {
		enq(n[0]+n[2], n[1]+n[3], n[2], n[3])
		return
	}
	switch v {
	case '|':
		enq(n[0], n[1]-1, 0, -1)
		enq(n[0], n[1]+1, 0, 1)
	case '-':
		enq(n[0]-1, n[1], -1, 0)
		enq(n[0]+1, n[1], 1, 0)
	case '/':
		switch {
		case n[3] == 1:
			enq(n[0]-1, n[1], -1, 0)
		case n[3] == -1:
			enq(n[0]+1, n[1], 1, 0)
		case n[2] == 1:
			enq(n[0], n[1]-1, 0, -1)
		case n[2] == -1:
			enq(n[0], n[1]+1, 0, 1)
		}
	case '\\':
		switch {
		case n[3] == 1:
			enq(n[0]+1, n[1], 1, 0)
		case n[3] == -1:
			enq(n[0]-1, n[1], -1, 0)
		case n[2] == 1:
			enq(n[0], n[1]+1, 0, 1)
		case n[2] == -1:
			enq(n[0], n[1]-1, 0, -1)
		}
	}
}

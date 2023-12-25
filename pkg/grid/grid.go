package grid

import (
	"fmt"
)

type Grid[T any] struct {
	ValToStr func(T, [2]int) string
	rotation int
	data     []T
	width    int
	height   int
}

func New[T any](w, h int, valToStr func(T, [2]int) string) *Grid[T] {
	return &Grid[T]{
		width:    w,
		height:   h,
		data:     make([]T, w*h),
		ValToStr: valToStr,
	}
}

func NewFromSlice[T any](lines [][]T, valToStr func(T, [2]int) string) *Grid[T] {
	g := &Grid[T]{
		width:    len(lines[0]),
		height:   len(lines),
		data:     make([]T, len(lines)*len(lines[0])),
		ValToStr: valToStr,
	}
	for y, line := range lines {
		for x, char := range line {
			g.Set(x, y, char)
		}
	}
	return g
}

func NewFromInput(input []string) *Grid[rune] {
	g := New[rune](len(input[0]), len(input), func(r rune, _ [2]int) string {
		return string(r)
	})
	for y, line := range input {
		for x, c := range line {
			g.Set(x, y, c)
		}
	}
	return g
}

func (g *Grid[T]) Get(x, y int) T {
	return g.data[g.mapIdx(x, y)]
}

func (g *Grid[T]) Set(x, y int, value T) {
	g.data[g.mapIdx(x, y)] = value
}

func (g *Grid[T]) mapIdx(x, y int) int {
	x2, y2 := x, y
	switch g.rotation {
	case 90:
		x2, y2 = g.width-y-1, x
	case 180:
		x2, y2 = g.width-x-1, g.height-y-1
	case 270:
		x2, y2 = y, g.height-x-1
	}
	return y2*g.width + x2
}

func (g *Grid[T]) Width() int {
	if g.rotation == 0 || g.rotation == 180 {
		return g.width
	}
	return g.height
}

func (g *Grid[T]) Height() int {
	if g.rotation == 0 || g.rotation == 180 {
		return g.height
	}
	return g.width
}

func (g *Grid[T]) Copy() *Grid[T] {
	var newDdata []T
	newDdata = append(newDdata, g.data...)
	return &Grid[T]{
		rotation: g.rotation,
		data:     newDdata,
		width:    g.width,
		height:   g.height,
		ValToStr: g.ValToStr,
	}
}

func (g *Grid[T]) IsInGrid(x, y int) bool {
	return x >= 0 && x < g.Width() && y >= 0 && y < g.Height()
}

func (g *Grid[T]) Col(i int) []T {
	col := make([]T, g.Height())
	for y := 0; y < g.Height(); y++ {
		col[y] = g.Get(i, y)
	}
	return col
}

func (g *Grid[T]) Row(i int) []T {
	row := make([]T, g.Width())
	for x := 0; x < g.Width(); x++ {
		row[x] = g.Get(x, i)
	}
	return row
}

func (g *Grid[T]) SubGrid(x1, y1, x2, y2 int) *Grid[T] {
	sub := New[T](x2-x1, y2-y1, g.ValToStr)
	for y := y1; y < y2; y++ {
		for x := x1; x < x2; x++ {
			sub.Set(x-x1, y-y1, g.Get(x, y))
		}
	}
	return sub
}

type Neighbor struct {
	Position [2]int
	Vector   [2]int
}

func (g *Grid[T]) Neighbors(pos [2]int, includeDiagonals bool) []Neighbor { // TODO: improve
	var neighbors []Neighbor
	vects := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	if includeDiagonals {
		vects = append(vects, [][2]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}}...)
	}
	for _, d := range vects {
		nP := [2]int{pos[0] + d[0], pos[1] + d[1]}
		if g.IsInGrid(nP[0], nP[1]) {
			neighbors = append(neighbors, Neighbor{Position: nP, Vector: d})
		}
	}
	return neighbors
}

// Rotate rotates the grid 90 degrees anti-clockwise N times.
func (g *Grid[T]) Rotate(times int) {
	v := g.rotation + times*90
	g.rotation = (v%360 + 360) % 360
}

type EnumerateItem[T any] struct {
	Position [2]int
	Value    T
}

func (g *Grid[T]) Enumerate() []EnumerateItem[T] {
	values := make([]EnumerateItem[T], len(g.data))
	i := 0
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			values[i] = EnumerateItem[T]{Position: [2]int{x, y}, Value: g.Get(x, y)}
			i++
		}
	}
	return values
}

func (g *Grid[T]) ToMap() map[[2]int]T {
	m := make(map[[2]int]T)
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			m[[2]int{x, y}] = g.Get(x, y)
		}
	}
	return m
}

func (g *Grid[T]) Hash() string {
	var s string
	for _, e := range g.Enumerate() {
		s += g.ValToStr(e.Value, e.Position) + "|"
	}
	return s
}

func (g *Grid[T]) Print() {
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			fmt.Print(g.ValToStr(g.Get(x, y), [2]int{x, y}))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (g *Grid[T]) Debug() {
	g.Print()
	fmt.Printf("Width: %d\n", g.Width())
	fmt.Printf("Height: %d\n", g.Height())
	fmt.Printf("Rotation: %d\n", g.rotation)
	fmt.Printf("\n")
}

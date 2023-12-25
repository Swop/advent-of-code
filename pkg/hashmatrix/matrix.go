package hashmatrix

import (
	"encoding/json"
	"fmt"
	"math"
	"sync"

	"github.com/Swop/advent-of-code/pkg/maps"
	math2 "github.com/Swop/advent-of-code/pkg/math"
)

func New[K Pos3D | Pos2D, V any]() *Matrix[K, V] {
	mins := make([]int, 3)
	maxs := make([]int, 3)
	for i := 0; i < 3; i++ {
		mins[i] = math.MaxInt
		maxs[i] = math.MinInt
	}
	return &Matrix[K, V]{
		elems: map[K]V{},
		mins:  mins,
		maxs:  maxs,
	}
}

type Matrix[K Pos3D | Pos2D, V any] struct {
	elems map[K]V
	mins  []int
	maxs  []int
	mutex sync.RWMutex
}

func (m *Matrix[K, V]) Get(k K) (V, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.getWithoutLocking(k)
}

func (m *Matrix[K, V]) getWithoutLocking(k K) (V, bool) {
	el, ok := m.elems[k]
	return el, ok
}

func (m *Matrix[K, V]) Set(k K, v V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.setWithoutLocking(k, v)
}

func getPosFromKey[K Pos2D | Pos3D](k K) Position {
	switch p := any(k).(type) {
	case Pos2D:
		return p
	case Pos3D:
		return p
	default:
		panic("unknown position type")
	}
}

func (m *Matrix[K, V]) setWithoutLocking(k K, v V) {
	m.elems[k] = v

	for i, ki := range getPosFromKey(k).Values() {
		if ki < m.mins[i] {
			m.mins[i] = ki
		}
		if ki > m.maxs[i] {
			m.maxs[i] = ki
		}
	}
}

func (m *Matrix[K, V]) SetIfNotPresent(k K, v V) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var ok bool
	if _, ok = m.getWithoutLocking(k); !ok {
		m.setWithoutLocking(k, v)
	}

	return !ok
}

func (m *Matrix[K, V]) UnSet(k K) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.elems, k)

	pos := getPosFromKey(k)
	maxes := make([]int, pos.Dimension())
	mins := make([]int, pos.Dimension())

	for i := 0; i < pos.Dimension(); i++ {
		mins[i] = math.MaxInt
		maxes[i] = math.MinInt
	}

	for elemKey := range m.elems {
		for i, ki := range getPosFromKey(elemKey).Values() {
			if ki < mins[i] {
				mins[i] = ki
			}
			if ki > maxes[i] {
				maxes[i] = ki
			}
		}
	}

	m.mins = mins
	m.maxs = maxes
}

type Neighbor[K Pos2D | Pos3D, V any] struct {
	Position K
	Vector   K
	Value    *V
}

func (m *Matrix[K, V]) Neighbors(k K, includeDiagonals bool) []Neighbor[K, V] {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	var neighbors []Neighbor[K, V]

	pos := getPosFromKey(k)
	kval := pos.Values()
	dimCount := pos.Dimension()

	var loop func(int, [3]int, bool)
	var nb func(int, [3]int, bool)
	loop = func(currentDim int, currentDimsValues [3]int, outOfMatrix bool) {
		for i := 0; i <= 2; i++ {
			currentDimsValues[dimCount-1-currentDim] = i
			if currentDim == dimCount-1 {
				nb(currentDim, currentDimsValues, outOfMatrix)
				continue
			}
			loop(currentDim+1, currentDimsValues, outOfMatrix)
		}
	}
	nb = func(currentDim int, currentDimsValues [3]int, outOfMatrix bool) {
		for i, d := range currentDimsValues {
			if i >= dimCount {
				break
			}
			outOfMatrix = outOfMatrix || kval[i]+d-1 < m.mins[i] || kval[i]+d-1 > m.maxs[i]
			if outOfMatrix {
				return
			}
		}

		var nP K
		var nV K
		delta := make([]int, dimCount)
		switch p := any(k).(type) {
		case Pos2D:
			nP = any(Pos2D{currentDimsValues[0], currentDimsValues[1]}).(K)
			delta[0] = currentDimsValues[0] - p.X
			delta[1] = currentDimsValues[1] - p.Y
			nV = any(Pos2D{delta[0], delta[1]}).(K)
		case Pos3D:
			nP = any(Pos3D{currentDimsValues[0], currentDimsValues[1], currentDimsValues[2]}).(K)
			delta[0] = currentDimsValues[0] - p.X
			delta[1] = currentDimsValues[1] - p.Y
			delta[2] = currentDimsValues[2] - p.Z
			nV = any(Pos3D{delta[0], delta[1], delta[2]}).(K)
		default:
			panic("unknown position type")
		}
		zeroCount := 0
		for _, d := range delta {
			if d == 0 {
				zeroCount++
			}
		}
		if zeroCount == dimCount || (dimCount-zeroCount > 1 && !includeDiagonals) {
			// same tile or diagonal
			return
		}
		var nVal *V
		if v, ok := m.elems[nP]; ok {
			nVal = &v
		}
		neighbors = append(neighbors, Neighbor[K, V]{Position: nP, Value: nVal, Vector: nV})
	}
	loop(0, [3]int{}, false)
	return neighbors
}

func (m *Matrix[K, V]) Copy() *Matrix[K, V] {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	m2 := New[K, V]()
	for k, v := range m.elems {
		m2.Set(k, v)
	}
	return m2
}

func (m *Matrix[K, V]) Size() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return len(maps.MapKeys(m.elems))
}

func (m *Matrix[K, V]) GetSubMap() map[K]V {
	return m.elems
}

func (m *Matrix[K, V]) Min(dimension int) int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.mins[dimension]
}

func (m *Matrix[K, V]) Max(dimension int) int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return m.maxs[dimension]
}

func (m *Matrix[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Mins  []int   `json:"mins"`
		Maxs  []int   `json:"maxs"`
		Elems map[K]V `json:"elems"`
	}{
		Mins:  m.mins,
		Maxs:  m.maxs,
		Elems: m.elems,
	})
}

func (m *Matrix[K, V]) Print2D() {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for y := m.mins[1]; y <= m.maxs[1]; y++ {
		for x := m.mins[0]; x <= m.maxs[0]; x++ {
			if v, ok := m.elems[any(Pos2D{x, y}).(K)]; ok {
				fmt.Print(v)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println("mins", m.mins)
	fmt.Println("maxs", m.maxs)
}

func (m *Matrix[K, V]) Print3D(axes [3]int, flip [3]bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	axisLabel := func(axis int) string {
		switch axis {
		case 0:
			return "x"
		case 1:
			return "y"
		default:
			return "z"
		}
	}
	fmt.Println("abs:", axisLabel(axes[0]), "- ord:", axisLabel(axes[1]))

	type loopSetup struct {
		start func(int) int
		cond  func(int, int) bool
		inc   func(int) int
	}
	var loopSetups [3]loopSetup
	for i := 0; i < 3; i++ {
		if flip[i] {
			loopSetups[i] = loopSetup{
				start: func(i int) int { return m.maxs[i] },
				cond:  func(i int, x int) bool { return x >= m.mins[i] },
				inc:   func(x int) int { return x - 1 },
			}
		} else {
			loopSetups[i] = loopSetup{
				start: func(i int) int { return m.mins[i] },
				cond:  func(i int, x int) bool { return x <= m.maxs[i] },
				inc:   func(x int) int { return x + 1 },
			}
		}
	}
	seen := map[[2]int]struct{}{}
	var mapping [3]int
	for y := loopSetups[1].start(axes[1]); loopSetups[1].cond(axes[1], y); y = loopSetups[1].inc(y) {
		for x := loopSetups[0].start(axes[0]); loopSetups[0].cond(axes[0], x); x = loopSetups[0].inc(x) {
			for z := loopSetups[2].start(axes[2]); loopSetups[2].cond(axes[2], z); z = loopSetups[2].inc(z) {
				mapping[axes[0]] = x
				mapping[axes[1]] = y
				mapping[axes[2]] = z
				v, valOk := m.elems[any(Pos3D{X: mapping[0], Y: mapping[1], Z: mapping[2]}).(K)]
				_, seenOk := seen[[2]int{x, y}]
				if !seenOk && valOk {
					fmt.Print(v)
					seen[[2]int{x, y}] = struct{}{}
				}
			}
			if _, seenOk := seen[[2]int{x, y}]; !seenOk {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println("mins", m.mins)
	fmt.Println("maxs", m.maxs)
}

func (m *Matrix[K, V]) Expand2D(dir [2]int) { // TODO adapt for N dimensions
	m.mutex.Lock()
	defer m.mutex.Unlock()

	mins := [2]int{m.mins[0], m.mins[1]}
	maxs := [2]int{m.maxs[0], m.maxs[1]}
	w := maxs[0] - mins[0] + 1
	h := maxs[1] - mins[1] + 1
	exp := func(dir [2]int) {
		if dir[0] == 0 && dir[1] == 0 {
			return
		}
		for y := mins[1]; y <= maxs[1]; y++ {
			for x := mins[0]; x <= maxs[0]; x++ {
				if v, ok := m.elems[any(Pos2D{x, y}).(K)]; ok {
					m.elems[any(Pos2D{x + dir[0]*w, y + dir[1]*h}).(K)] = v
				}
			}
		}
	}

	exp([2]int{dir[0], 0})
	exp([2]int{0, dir[1]})
	if dir[0] != 0 && dir[1] != 0 {
		exp(dir)
	}

	switch {
	case dir[0] == -1:
		m.mins[0] -= w
	case dir[0] == 1:
		m.maxs[0] += w
	}
	switch {
	case dir[1] == -1:
		m.mins[1] -= h
	case dir[1] == 1:
		m.maxs[1] += h
	}
}

type Position interface {
	Dimension() int
	Values() []int
}

type Pos2D struct {
	X, Y int
}

func (p2d Pos2D) String() string {
	return fmt.Sprintf("%d|%d", p2d.X, p2d.Y)
}

func (p2d Pos2D) Dimension() int {
	return 2
}

func (p2d Pos2D) Values() []int {
	return []int{p2d.X, p2d.Y}
}

func (p2d Pos2D) ManhattanDistance(other Pos2D) int {
	return math2.ManhattanDistance(p2d.X, p2d.Y, other.X, other.Y)
}

type Pos3D struct {
	X, Y, Z int
}

func (p3d Pos3D) Dimension() int {
	return 3
}

func (p3d Pos3D) String() string {
	return fmt.Sprintf("%d|%d|%d", p3d.X, p3d.Y, p3d.Z)
}

func (p3d Pos3D) Values() []int {
	return []int{p3d.X, p3d.Y, p3d.Z}
}

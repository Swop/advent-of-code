package heap

import (
	"sync"
)

func New[T any](comp func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		comp: comp,
	}
}

type Heap[T any] struct {
	items []T
	comp  func(a, b T) bool
	mutex sync.RWMutex
}

func (h *Heap[T]) Size() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	return h.unsafeSize()
}

func (h *Heap[T]) unsafeSize() int {
	return len(h.items)
}

func (h *Heap[T]) Push(item T) T {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.items = append(h.items, item)
	h.up(h.unsafeSize() - 1)

	return item
}

func (h *Heap[T]) Pop() T {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	n := h.unsafeSize() - 1
	if n > 0 {
		h.swap(0, n)
		h.down()
	}

	v := h.items[n]
	h.items = h.items[0:n]

	return v
}

func (h *Heap[T]) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Heap[T]) up(jj int) {
	for {
		i := heapParent(jj)
		if i == jj || !h.comp(h.items[jj], h.items[i]) {
			break
		}
		h.swap(i, jj)
		jj = i
	}
}

func (h *Heap[T]) down() {
	n := h.unsafeSize() - 1
	i1 := 0
	for {
		j1 := heapLeft(i1)
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		j2 := heapRight(i1)
		if j2 < n && h.comp(h.items[j2], h.items[j1]) {
			j = j2
		}
		if !h.comp(h.items[j], h.items[i1]) {
			break
		}
		h.swap(i1, j)
		i1 = j
	}
}

func heapParent(i int) int {
	return (i - 1) / 2
}

func heapLeft(i int) int {
	return (i * 2) + 1
}

func heapRight(i int) int {
	return heapLeft(i) + 1
}

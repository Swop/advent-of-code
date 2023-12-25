package linkedlist

import "fmt"

type Node[T any] struct {
	data T
	prev *Node[T]
	next *Node[T]
}

type DoubleLinkedList[T any] struct {
	len  int
	tail *Node[T]
	head *Node[T]
}

func NewDoubleLinkedList[T any]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{}
}

func (d *DoubleLinkedList[T]) AddFront(data T) {
	n := &Node[T]{
		data: data,
	}
	if d.head == nil {
		d.head = n
		d.tail = n
	} else {
		n.next = d.head
		d.head.prev = n
		d.head = n
	}
	d.len++
}

func (d *DoubleLinkedList[T]) AddEnd(data T) {
	n := &Node[T]{
		data: data,
	}
	if d.head == nil {
		d.head = n
		d.tail = n
	} else {
		n.prev = d.tail
		d.tail.next = n
		d.tail = n
	}
	d.len++
}

func (d *DoubleLinkedList[T]) TraverseForward(fn func(T)) {
	d.traverse(1, func(n *Node[T]) {
		fn(n.data)
	})
}

func (d *DoubleLinkedList[T]) TraverseBackward(fn func(T)) {
	d.traverse(-1, func(n *Node[T]) {
		fn(n.data)
	})
}

func (d *DoubleLinkedList[T]) Filter(fn func(T) bool) {
	d.traverse(1, func(n *Node[T]) {
		if !fn(n.data) {
			if n.prev == nil {
				d.head = n.next
			} else {
				n.prev.next = n.next
			}
			if n.next == nil {
				d.tail = n.prev
			} else {
				n.next.prev = n.prev
			}
		}
	})
}

// Replace replaces the first element that matches the predicate with the given data.
// If no element matches the predicate (and the search hasn't been stopped by the predicate), the element is added
// to the end of the list.
// The given function should return two booleans, the first one indicating whether the element should be replaced,
// the second one indicating whether the search should be stopped.
func (d *DoubleLinkedList[T]) Replace(fn func(T) (bool, bool), data T) {
	replaced := false
	d.traverse(1, func(n *Node[T]) {
		matched, stop := fn(n.data) //nolint:staticcheck
		if matched {
			n.data = data
			replaced = true
			if stop {
				return
			}
		}
	})
	if !replaced {
		d.AddEnd(data)
	}
}

func (d *DoubleLinkedList[T]) Size() int {
	return d.len
}

func (d *DoubleLinkedList[T]) String() string {
	var s string
	d.traverse(1, func(n *Node[T]) {
		s += fmt.Sprintf("%v", n.data)
		if n.next != nil {
			s += " "
		}
	})
	return s
}

func (d *DoubleLinkedList[T]) traverse(dir int, fn func(*Node[T])) {
	if d.head == nil {
		return
	}
	var c *Node[T]
	c = d.head
	if dir < 0 {
		c = d.tail
	}
	for c != nil {
		fn(c)
		c = c.next
		if dir < 0 {
			c = c.prev
		}
	}
}

package hashmap

import "github.com/Swop/advent-of-code/pkg/linkedlist"

type HashFunc[H any, K comparable] func(H) K

type HashMap[H any, K comparable, V any] struct {
	data map[K]*linkedlist.DoubleLinkedList[V]
	hFn  HashFunc[H, K]
}

func New[H any, K comparable, V any](hFn HashFunc[H, K]) *HashMap[H, K, V] {
	return &HashMap[H, K, V]{
		data: map[K]*linkedlist.DoubleLinkedList[V]{},
		hFn:  hFn,
	}
}

func (hm *HashMap[H, K, V]) Size() int {
	return len(hm.data)
}

func (hm *HashMap[H, K, V]) Keys() []K {
	keys := make([]K, 0, len(hm.data))
	for k := range hm.data {
		keys = append(keys, k)
	}
	return keys
}

func (hm *HashMap[H, K, V]) Values() []*linkedlist.DoubleLinkedList[V] {
	values := make([]*linkedlist.DoubleLinkedList[V], 0, len(hm.data))
	for _, v := range hm.data {
		values = append(values, v)
	}
	return values
}

func (hm *HashMap[H, K, V]) Has(keyInput H) bool {
	_, ok := hm.data[hm.hFn(keyInput)]
	return ok
}

func (hm *HashMap[H, K, V]) Enumerate() map[K]*linkedlist.DoubleLinkedList[V] {
	return hm.data
}

func (hm *HashMap[H, K, V]) Get(keyInput H) (*linkedlist.DoubleLinkedList[V], bool) {
	v, ok := hm.data[hm.hFn(keyInput)]
	return v, ok
}

func (hm *HashMap[H, K, V]) SetWithAdd(keyInput H, value V) {
	h := hm.hFn(keyInput)
	l, ok := hm.data[h]
	if !ok {
		l = linkedlist.NewDoubleLinkedList[V]()
		hm.data[h] = l
	}
	l.AddEnd(value)
}

func (hm *HashMap[H, K, V]) SetWithReplace(keyInput H, fn func(V) (bool, bool), value V) {
	h := hm.hFn(keyInput)
	l, ok := hm.data[h]
	if !ok {
		l = linkedlist.NewDoubleLinkedList[V]()
		hm.data[h] = l
	}
	l.Replace(fn, value)
}

func (hm *HashMap[H, K, V]) Unset(keyInput H, fn func(V) bool) {
	h := hm.hFn(keyInput)
	l, ok := hm.data[h]
	if !ok {
		return
	}
	l.Filter(fn)
	if l.Size() == 0 {
		delete(hm.data, h)
	}
}

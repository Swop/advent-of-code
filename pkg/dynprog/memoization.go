package dynprog

type (
	MemoizedFunc[V any]            func(...any) V
	MemoizedHashFunc[K comparable] func(...any) K
)

func Memoize[K comparable, V any](hashFunc MemoizedHashFunc[K], initialFunc MemoizedFunc[V]) MemoizedFunc[V] {
	cache := map[K]V{}
	return func(args ...any) V {
		h := hashFunc(args...)
		if v, ok := cache[h]; ok {
			return v
		}
		v := initialFunc(args...)
		cache[h] = v
		return v
	}
}

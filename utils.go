package aoc

type Set[T comparable] map[T]bool

func (s Set[T]) Add(e T) {
	s[e] = true
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) toArray() []T {
	arr := make([]T, 0)
	for e := range s {
		arr = append(arr, e)
	}
	return arr
}

package aoc

import (
	"bufio"
	"os"
)

type Set[T comparable] map[T]bool

func (s Set[T]) Add(e T) {
	s[e] = true
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Has(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set[T]) ToArray() []T {
	arr := make([]T, 0)
	for e := range s {
		arr = append(arr, e)
	}
	return arr
}

func (s Set[T]) Remove(e T) {
	delete(s, e)
}

func (s Set[T]) FromArray(arr []T) {
	for _, e := range arr {
		s.Add(e)
	}
}

func (s Set[T]) Extend(other Set[T]) {
	for e := range other {
		s.Add(e)
	}
}

func LoadInputLines(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return n
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func Copy2DArray[T any](arr [][]T) [][]T {
	newArr := make([][]T, len(arr))
	for i, row := range arr {
		newArr[i] = make([]T, len(row))
		copy(newArr[i], row)
	}
	return newArr
}

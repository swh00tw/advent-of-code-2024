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

func (s Set[T]) ToArray() []T {
	arr := make([]T, 0)
	for e := range s {
		arr = append(arr, e)
	}
	return arr
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

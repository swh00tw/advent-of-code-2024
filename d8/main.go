package main

import (
	"bufio"
	"fmt"
	"os"
)

type Loc struct {
	X int
	Y int
}

func (l *Loc) Add(other Loc) {
	l.X += other.X
	l.Y += other.Y
}

func (n *Loc) String() string {
	return fmt.Sprintf("%d,%d", n.X, n.Y)
}

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

func (s Set[T]) Extend(other Set[T]) {
	for e := range other {
		s.Add(e)
	}
}

type Map struct {
	M      int
	N      int
	freq   []byte
	matrix [][]byte
}

func (m *Map) InBound(loc Loc) bool {
	x := loc.X
	y := loc.Y
	return x >= 0 && x < m.M && y >= 0 && y < m.N
}

func (m *Map) FindAntiNode(a, b Loc) Set[Loc] {
	ans := Set[Loc]{}
	first := Loc{
		X: 2*a.X - b.X,
		Y: 2*a.Y - b.Y,
	}
	if first.X >= 0 && first.Y >= 0 && first.X < m.M && first.Y < m.N {
		ans.Add(first)
	}
	second := Loc{
		X: 2*b.X - a.X,
		Y: 2*b.Y - a.Y,
	}
	if second.X >= 0 && second.Y >= 0 && second.X < m.M && second.Y < m.N {
		ans.Add(second)
	}
	return ans
}

func (m *Map) FindAntiNode2(a, b Loc) Set[Loc] {
	ans := Set[Loc]{}
	ans.Add(a)
	ans.Add(b)
	diff := Loc{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
	negativeDiff := Loc{
		X: -diff.X,
		Y: -diff.Y,
	}
	curr := Loc{
		X: a.X,
		Y: a.Y,
	}
	for m.InBound(curr) {
		ans.Add(curr)
		curr.Add(diff)
	}
	curr = Loc{
		X: a.X,
		Y: a.Y,
	}
	for m.InBound(curr) {
		ans.Add(curr)
		curr.Add(negativeDiff)
	}
	return ans
}

func (m *Map) FindAntinodes(freq byte, findAntiNodeFunction func(a, b Loc) Set[Loc]) Set[Loc] {
	antennasLocs := make([]Loc, 0)
	for i, row := range m.matrix {
		for j, cell := range row {
			if cell == freq {
				loc := Loc{i, j}
				antennasLocs = append(antennasLocs, loc)
			}
		}
	}

	antinodeLocs := Set[Loc]{}
	// for any pair
	for i, locA := range antennasLocs {
		for j, locB := range antennasLocs {
			if j <= i {
				continue
			}
			locs := findAntiNodeFunction(locA, locB)
			antinodeLocs.Extend(locs)
		}
	}
	return antinodeLocs
}

func part1(m Map) int {
	locsSet := Set[Loc]{}
	for _, freq := range m.freq {
		locs := m.FindAntinodes(freq, m.FindAntiNode)
		//fmt.Println(string(freq), locs)
		locsSet.Extend(locs)
	}
	return locsSet.Len()
}

func part2(m Map) int {
	locsSet := Set[Loc]{}
	for _, freq := range m.freq {
		locs := m.FindAntinodes(freq, m.FindAntiNode2)
		//fmt.Println(string(freq), locs)
		locsSet.Extend(locs)
	}
	return locsSet.Len()
}

func loadInput() Map {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	freqSet := Set[byte]{}
	matrix := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []byte(line)
		matrix = append(matrix, row)
		for _, b := range row {
			if b == '.' {
				continue
			}
			freqSet.Add(b)
		}
	}
	m := len(matrix)
	n := len(matrix[0])
	return Map{
		M:      m,
		N:      n,
		matrix: matrix,
		freq:   freqSet.toArray(),
	}
}

func main() {
	m := loadInput()

	fmt.Println(part1(m))
	fmt.Println(part2(m))
}

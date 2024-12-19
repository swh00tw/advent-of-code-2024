package main

import (
	"fmt"
	"github.com/swh00tw/aoc"
	"strings"
)

var filename = "input.txt"

func loadInput() (aoc.Set[string], []string) {
	lines := aoc.LoadInputLines(filename)

	towels := make(aoc.Set[string])
	designs := []string{}
	for i, line := range lines {
		if i == 0 {
			strs := strings.Split(line, ", ")
			towels.FromArray(strs)
		}
		if i == 1 {
			continue
		}
		designs = append(designs, line)
	}
	return towels, designs
}

func part1(towels aoc.Set[string], designs []string) int {
	cache := make(map[string]bool)

	var isPossible func(i, j int) bool
	isPossible = func(i, j int) bool {
		key := fmt.Sprintf("%d,%d", i, j)
		if v, ok := cache[key]; ok {
			return v
		}
		s := designs[i][j:]
		if len(s) == 0 {
			return true
		}
		ans := false
		n := len(s)
		for l := n; l >= 1; l-- {
			prefix := s[:l]
			if towels.Has(prefix) {
				if isPossible(i, j+l) {
					ans = true
					break
				}
			}
		}
		cache[key] = ans
		return ans
	}

	count := 0
	for i, _ := range designs {
		if isPossible(i, 0) {
			count++
		}
	}
	return count
}

func part2(towels aoc.Set[string], designs []string) int {
	cache := make(map[string]int)

	var isPossible func(i, j int) int
	isPossible = func(i, j int) int {
		key := fmt.Sprintf("%d,%d", i, j)
		if v, ok := cache[key]; ok {
			return v
		}
		s := designs[i][j:]
		if len(s) == 0 {
			return 1
		}
		ans := 0
		n := len(s)
		for l := n; l >= 1; l-- {
			prefix := s[:l]
			if towels.Has(prefix) {
				subans := isPossible(i, j+l)
				if subans > 0 {
					ans += subans
				}
			}
		}
		cache[key] = ans
		return ans
	}

	count := 0
	for i, _ := range designs {
		count += isPossible(i, 0)
	}
	return count
}

func main() {
	towels, designs := loadInput()

	fmt.Println(part1(towels, designs))
	fmt.Println(part2(towels, designs))
}

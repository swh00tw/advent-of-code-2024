package main

import (
	"bufio"
	"fmt"
	"github.com/swh00tw/aoc"
	"os"
	"strconv"
)

func pos2Key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func loadInput() [][]int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, b := range line {
			v, _ := strconv.Atoi(string(b))
			row = append(row, v)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func calculateScore(matrix [][]int, x, y int) int {
	m := len(matrix)
	n := len(matrix[0])
	ninePosition := aoc.Set[string]{}
	deltas := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	var findNext func(x, y int, target int)
	findNext = func(x, y int, target int) {
		// edge case: out of bound
		if x < 0 || y < 0 || x >= m || y >= n {
			return
		}
		// edge case, if not match
		if matrix[x][y] != target {
			return
		}
		// base case
		if target == 9 {
			// add position to set and return
			ninePosition.Add(pos2Key(x, y))
			return
		}
		// general case
		for _, d := range deltas {
			nx := x + d[0]
			ny := y + d[1]
			findNext(nx, ny, target+1)
		}
	}

	findNext(x, y, 0)
	return ninePosition.Len()
}

func part1(matrix [][]int) int {
	ans := 0
	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				ans += calculateScore(matrix, i, j)
			}
		}
	}
	return ans
}

func poss2Key(pos [][]int) string {
	key := ""
	for _, p := range pos {
		key += pos2Key(p[0], p[1])
	}
	return key
}

func calculateRating(matrix [][]int, x, y int) int {
	m := len(matrix)
	n := len(matrix[0])
	ninePaths := aoc.Set[string]{}
	deltas := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	var findNext func(x, y int, target int, path [][]int)
	findNext = func(x, y int, target int, path [][]int) {
		nextPath := append(path, []int{x, y})
		// edge case: out of bound
		if x < 0 || y < 0 || x >= m || y >= n {
			return
		}
		// edge case, if not match
		if matrix[x][y] != target {
			return
		}
		// base case
		if target == 9 {
			// add position to set and return
			ninePaths.Add(poss2Key(nextPath))
			return
		}
		// general case
		for _, d := range deltas {
			nx := x + d[0]
			ny := y + d[1]
			findNext(nx, ny, target+1, nextPath)
		}
	}

	findNext(x, y, 0, [][]int{})
	return ninePaths.Len()
}

func part2(matrix [][]int) int {
	ans := 0
	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				ans += calculateRating(matrix, i, j)
			}
		}
	}
	return ans
}

func main() {
	matrix := loadInput()

	fmt.Println(part1(matrix))
	fmt.Println(part2(matrix))
}

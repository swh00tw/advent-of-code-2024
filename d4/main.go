package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadInput() ([][]byte, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	mat := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []byte{}
		for _, c := range line {
			row = append(row, byte(c))
		}
		mat = append(mat, row)
	}
	return mat, scanner.Err()
}

func part1(mat [][]byte) int {
	/*
	* DFS, backtracing
	 */
	m := len(mat)
	n := len(mat[0])
	ans := 0
	direction := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	find := func(pos []int, target string, dir []int) {
		curr := pos
		bytes := []byte{}
		for curr[0] >= 0 && curr[0] < m && curr[1] >= 0 && curr[1] < n && len(bytes) < len(target) {
			bytes = append(bytes, mat[curr[0]][curr[1]])
			curr[0] += dir[0]
			curr[1] += dir[1]
		}
		if string(bytes) == target {
			ans++
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, dir := range direction {
				find([]int{i, j}, "XMAS", dir)
			}
		}
	}

	return ans
}

func part2(input [][]byte) int {
	ans := 0
	m := len(input)
	n := len(input[0])

	isXmas := func(pos []int) bool {
		if pos[0] < 1 || pos[0] >= m-1 || pos[1] < 1 || pos[1] >= n-1 {
			return false
		}
		if input[pos[0]][pos[1]] != 'A' {
			return false
		}
		left := []byte{
			input[pos[0]-1][pos[1]-1],
			input[pos[0]+1][pos[1]+1],
		}
		leftOk := string(left) == "MS" || string(left) == "SM"
		right := []byte{
			input[pos[0]-1][pos[1]+1],
			input[pos[0]+1][pos[1]-1],
		}
		rightOk := string(right) == "MS" || string(right) == "SM"

		if !rightOk || !leftOk {
			return false
		}
		return true
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isXmas([]int{i, j}) {
				ans++
			}
		}
	}

	return ans
}

func main() {
	mat, _ := loadInput()

	fmt.Println("Part 1: ", part1(mat))
	fmt.Println("Part 2: ", part2(mat))
}

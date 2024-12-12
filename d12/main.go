package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadInput() [][]byte {
	file, _ := os.Open("input.txt")
	defer file.Close()

	matrix := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []byte(line))
	}
	return matrix
}

func findAreaPrice(matrix [][]byte, visited *[][]bool, i, j int) int64 {
	//fmt.Println("pos: ", i, j, ", char: ", string(matrix[i][j]))
	// BFS
	deltas := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}

	area := 0
	peri := 0

	queue := make([][]int, 0)
	queue = append(queue, []int{i, j})
	for len(queue) > 0 {
		// pop
		cell := queue[0]
		queue = queue[1:]
		x, y := cell[0], cell[1]
		if (*visited)[x][y] == true {
			continue
		}
		(*visited)[x][y] = true

		area++

		for _, d := range deltas {
			nx := x + d[0]
			ny := y + d[1]
			if nx >= 0 && nx < len(matrix) && ny >= 0 && ny < len(matrix[0]) && matrix[nx][ny] == matrix[x][y] {
				queue = append(queue, []int{nx, ny})
			} else {
				peri++
			}
		}
	}

	return int64(peri) * int64(area)
}

func part1(matrix [][]byte) int64 {
	m := len(matrix)
	n := len(matrix[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	ans := int64(0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				ans += findAreaPrice(matrix, &visited, i, j)
			}
		}
	}
	return ans
}

func findAreaPrice2(matrix [][]byte, visited *[][]bool, i, j int) int64 {
	//fmt.Println("pos: ", i, j, ", char: ", string(matrix[i][j]))
	// BFS
	deltas := [][]int{
		{-1, 0}, // UP
		{0, 1},  // RIGHT
		{1, 0},  // DOWN
		{0, -1}, // LEFT
	}

	area := 0
	corners := 0

	queue := make([][]int, 0)
	queue = append(queue, []int{i, j})
	for len(queue) > 0 {
		// pop
		cell := queue[0]
		queue = queue[1:]
		x, y := cell[0], cell[1]
		if (*visited)[x][y] == true {
			continue
		}
		(*visited)[x][y] = true

		area++

		res := []bool{}
		for _, d := range deltas {
			nx := x + d[0]
			ny := y + d[1]
			if nx >= 0 && nx < len(matrix) && ny >= 0 && ny < len(matrix[0]) && matrix[nx][ny] == matrix[x][y] {
				queue = append(queue, []int{nx, ny})
				res = append(res, true)
			} else {
				res = append(res, false)
			}
		}

		// count corners
		// inner corners
		if res[0] && res[1] {
			if x-1 >= 0 && y+1 < len(matrix[0]) && matrix[x-1][y+1] != matrix[x][y] {
				corners++
			}
		}
		if res[1] && res[2] {
			if x+1 < len(matrix) && y+1 < len(matrix[0]) && matrix[x+1][y+1] != matrix[x][y] {
				corners++
			}
		}
		if res[2] && res[3] {
			if x+1 < len(matrix) && y-1 >= 0 && matrix[x+1][y-1] != matrix[x][y] {
				corners++
			}
		}
		if res[3] && res[0] {
			if x-1 >= 0 && y-1 >= 0 && matrix[x-1][y-1] != matrix[x][y] {
				corners++
			}
		}
		// outer corners
		if !res[0] && !res[1] {
			corners++
		}
		if !res[1] && !res[2] {
			corners++
		}
		if !res[2] && !res[3] {
			corners++
		}
		if !res[3] && !res[0] {
			corners++
		}
	}

	fmt.Println("char: ", string(matrix[i][j]))
	fmt.Println("corners: ", corners)
	fmt.Println("area: ", area)
	return int64(area) * int64(corners)
}

func part2(matrix [][]byte) int64 {
	m := len(matrix)
	n := len(matrix[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	ans := int64(0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				ans += findAreaPrice2(matrix, &visited, i, j)
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

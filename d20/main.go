package main

import (
	"fmt"
	"github.com/swh00tw/aoc"
	"math"
)

var fileName = "input.txt"

func loadInput() [][]byte {
	lines := aoc.LoadInputLines(fileName)
	matrix := make([][]byte, len(lines))
	for i, line := range lines {
		matrix[i] = []byte(line)
	}
	return matrix
}

func BFS(x, y int, matrix [][]byte) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
		for j, _ := range visited[i] {
			visited[i][j] = math.MaxInt
		}
	}

	queue := make([][]int, 0)
	queue = append(queue, []int{x, y})
	visited[x][y] = 0

	for len(queue) > 0 {
		// pop
		curr := queue[0]
		queue = queue[1:]
		x, y := curr[0], curr[1]

		for _, dir := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && matrix[nx][ny] != '#' && visited[nx][ny] == math.MaxInt {
				visited[nx][ny] = visited[x][y] + 1
				queue = append(queue, []int{nx, ny})
			}
		}
	}

	return visited
}

func getStartEnd(matrix [][]byte) [][]int {
	start := make([]int, 2)
	end := make([]int, 2)
	for i, row := range matrix {
		for j, cell := range row {
			if cell == 'S' {
				start[0], start[1] = i, j
			} else if cell == 'E' {
				end[0], end[1] = i, j
			}
		}
	}
	return [][]int{start, end}
}

func solve(matrix [][]byte, threshold int, cheatAllowed int) int {
	startEnd := getStartEnd(matrix)
	start, _ := startEnd[0], startEnd[1]
	visited := BFS(start[0], start[1], matrix)
	m := len(matrix)
	n := len(matrix[0])

	save := make(map[int]aoc.Set[string])

	pathKey := func(x1, y1, x2, y2 int) string {
		return fmt.Sprintf("%d,%d,%d,%d", x1, y1, x2, y2)
	}

	// for each cell on the path
	for i, row := range matrix {
		for j, cell := range row {
			if cell == '.' || cell == 'S' {
				visitedCheat := aoc.Copy2DArray(visited)

				queue := make([][]int, 0)
				queue = append(queue, []int{i, j, 0}) // x, y, iter
				for len(queue) > 0 {
					// pop
					curr := queue[0]
					queue = queue[1:]
					x, y, iter := curr[0], curr[1], curr[2]
					if iter > 0 && (matrix[x][y] == 'E' || matrix[x][y] == '.') {
						diff := visited[x][y] - visitedCheat[x][y]
						if _, ok := save[diff]; !ok {
							save[diff] = aoc.Set[string]{}
						}
						key := pathKey(x, y, i, j)
						save[diff].Add(key)
					}
					if matrix[x][y] == 'E' {
						continue
					}

					for _, dir := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
						nx, ny := x+dir[0], y+dir[1]
						if nx >= 0 && nx < m && ny >= 0 && ny < n {
							if visitedCheat[x][y]+1 <= visitedCheat[nx][ny] && iter < cheatAllowed {
								visitedCheat[nx][ny] = visitedCheat[x][y] + 1
								queue = append(queue, []int{nx, ny, iter + 1})
							}
						}
					}
				}

			}
		}
	}

	ans := make(map[int]int)
	for k, v := range save {
		if k >= threshold {
			ans[k] = v.Len()
		}
	}
	fmt.Println("ans: ", ans)

	cnt := 0
	for _, v := range ans {
		cnt += v
	}

	return cnt
}

func main() {
	matrix := loadInput()

	fmt.Println("Part 1: ", solve(matrix, 1, 2))
	fmt.Println("Part 2: ", solve(matrix, 100, 20))
}

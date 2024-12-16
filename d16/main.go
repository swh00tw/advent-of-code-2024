package main

import (
	"fmt"
	"github.com/swh00tw/aoc"
	"math"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

var directionsToVec = map[Direction][]int{
	Up:    {-1, 0},
	Right: {0, 1},
	Down:  {1, 0},
	Left:  {0, -1},
}

func getNextDirection(direction Direction) []Direction {
	dirs := []Direction{}
	currVec := directionsToVec[direction]
	for k, v := range directionsToVec {
		if v[0] == -currVec[0] && v[1] == -currVec[1] {
			continue
		}
		dirs = append(dirs, k)
	}
	return dirs
}

func loadInput() [][]byte {
	maze := [][]byte{}
	lines := aoc.LoadInputLines("input.txt")
	for _, line := range lines {
		maze = append(maze, []byte(line))
	}
	return maze
}

func getStartAndEnd(maze [][]byte) [][]int {
	start := []int{}
	end := []int{}

	for i, line := range maze {
		for j, cell := range line {
			if cell == 'S' {
				start = append(start, i, j)
			}
			if cell == 'E' {
				end = append(end, i, j)
			}
		}
	}
	ans := [][]int{
		start, end,
	}
	return ans
}

type Visit struct {
	X         int
	Y         int
	Direction Direction
	Cost      int
}

func (v Visit) Pos() string {
	return fmt.Sprintf("%d,%d", v.X, v.Y)
}

func bfs(maze [][]byte, initVisit Visit, endPos []int) Visit {
	m := len(maze)
	n := len(maze[0])

	visit := make([][]Visit, m)
	for i := range visit {
		visit[i] = make([]Visit, n)
		for j := range visit[i] {
			visit[i][j] = Visit{
				i, j, Up, math.MaxInt,
			}
		}
	}

	// BFS
	queue := []Visit{}
	queue = append(queue, initVisit)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		// stop at end
		if cur.X == endPos[0] && cur.Y == endPos[1] {
			continue
		}

		nextDirections := getNextDirection(cur.Direction)
		for _, nextDirection := range nextDirections {
			vec := directionsToVec[nextDirection]
			nx := cur.X + vec[0]
			ny := cur.Y + vec[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && (maze[nx][ny] == '.' || maze[nx][ny] == 'E') {

				newCost := cur.Cost + 1
				if nextDirection != cur.Direction {
					newCost += 1000
				}

				nextVisit := Visit{
					nx, ny, nextDirection, newCost,
				}

				if newCost <= visit[nx][ny].Cost {
					queue = append(queue, nextVisit)
				}
				// update visit
				if visit[nx][ny].Cost > nextVisit.Cost {
					visit[nx][ny] = nextVisit
				}
			}
		}
	}

	return visit[endPos[0]][endPos[1]]
}

func part1(maze [][]byte) int {
	startEnd := getStartAndEnd(maze)
	startPos := startEnd[0]
	endPos := startEnd[1]

	// BFS
	dir := Right

	initVisit := Visit{
		startPos[0], startPos[1], dir, 0,
	}
	ans := bfs(maze, initVisit, endPos)

	return ans.Cost
}

func part2(maze [][]byte) int {
	startEnd := getStartAndEnd(maze)
	startPos := startEnd[0]
	endPos := startEnd[1]
	initVisit := Visit{
		startPos[0], startPos[1], Right, 0,
	}
	totalCost := bfs(maze, initVisit, endPos).Cost

	pos := aoc.Set[string]{}
	for i, line := range maze {
		for j, cell := range line {
			if cell == '#' {
				continue
			}
			if cell == '.' {
				state := bfs(maze, initVisit, []int{i, j})
				visit := bfs(maze, state, endPos)
				if visit.Cost == totalCost {
					pos.Add(state.Pos())
				}
			}
		}
	}
	return pos.Len() + 2
}

func main() {
	maze := loadInput()

	fmt.Printf("Part 1: %d\n", part1(maze))
	fmt.Printf("Part 2: %d\n", part2(maze))
}

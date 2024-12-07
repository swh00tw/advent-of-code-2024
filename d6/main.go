package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadInput() ([][]byte, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return nil, err
	}
	defer file.Close()

	matrix := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)
		matrix = append(matrix, bytes)
	}

	return matrix, nil
}

type Set[T comparable] map[T]bool

func (s *Set[T]) Add(value T) {
	(*s)[value] = true
}

func (s *Set[T]) Len() int {
	return len(*s)
}

type Coord struct {
	X int
	Y int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

func findGuardInitPos(matrix [][]byte) Coord {
	for x, row := range matrix {
		for y, cell := range row {
			if cell == '^' {
				return Coord{x, y}
			}
		}
	}
	return Coord{-1, -1}
}

// direction: 0 - up, 1 - right, 2 - down, 3 - left
func getNextDirection(currDir int) int {
	return (currDir + 1) % 4
}

func getNextPos(matrix [][]byte, curr Coord, direction int) (Coord, int) {
	nx := curr.X
	ny := curr.Y
	switch direction {
	case 0:
		nx--
	case 1:
		ny++
	case 2:
		nx++
	case 3:
		ny--
	}
	m := len(matrix)
	n := len(matrix[0])
	if nx < 0 || nx >= m || ny < 0 || ny >= n {
		return Coord{-1, -1}, 0
	}
	if matrix[nx][ny] == '#' {
		// turn right
		nextDirection := getNextDirection(direction)
		return getNextPos(matrix, curr, nextDirection)
	}
	return Coord{nx, ny}, direction
}

func part1(matrix [][]byte) int {
	guardPos := findGuardInitPos(matrix)
	direction := 0
	visited := Set[string]{}
	for guardPos.X != -1 {
		visited.Add(guardPos.String())
		nextPos, nextDir := getNextPos(matrix, guardPos, direction)
		guardPos = nextPos
		direction = nextDir
	}

	return visited.Len()
}

func part2experiment(matrix [][]byte) bool {
	m := len(matrix)
	n := len(matrix[0])
	guardPos := findGuardInitPos(matrix)
	direction := 0
	iterations := 0

	for guardPos.X != -1 {
		if iterations > m*n {
			return false
		}
		nextPos, nextDir := getNextPos(matrix, guardPos, direction)
		guardPos = nextPos
		direction = nextDir
		iterations++
	}
	return true
}

func part2(matrix [][]byte) int {
	ans := 0
	for i, row := range matrix {
		for j, cell := range row {
			if cell == '#' || cell == '^' {
				continue
			}
			matrix[i][j] = '#'
			// run experiment
			if !part2experiment(matrix) {
				ans++
			}
			matrix[i][j] = cell
		}
	}
	return ans
}

func main() {
	matrix, err := loadInput()
	if err != nil {
		fmt.Println("Error loading input")
		return
	}

	fmt.Println(part1(matrix))
	fmt.Println(part2(matrix))
}

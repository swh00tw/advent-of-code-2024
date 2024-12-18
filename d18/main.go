package main

import (
	"fmt"
	"github.com/swh00tw/aoc"
	"math"
	"strconv"
	"strings"
)

var inputFileName = "input.txt"
var gridSize = 71
var numBytes = 1024

type Point struct {
	X, Y int
}

func (p Point) ToString() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type Grid struct {
	points    []Point
	pointsSet aoc.Set[string]
	m         int
	n         int
}

func (g *Grid) Print() {
	for i := 0; i < g.m; i++ {
		for j := 0; j < g.n; j++ {
			if g.pointsSet.Has(Point{i, j}.ToString()) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (g *Grid) GetShortestPath(start Point, end Point) int {
	// BFS
	queue := []Point{start}
	visited := make([][]bool, g.m)
	steps := make([][]int, g.m)
	for i := 0; i < g.m; i++ {
		visited[i] = make([]bool, g.n)
		steps[i] = make([]int, g.n)
		for j := 0; j < g.n; j++ {
			steps[i][j] = math.MaxInt
		}
	}

	steps[start.X][start.Y] = 0
	visited[start.X][start.Y] = true
	for len(queue) > 0 {
		// pop
		p := queue[0]
		queue = queue[1:]
		if p.X == end.X && p.Y == end.Y {
			return steps[end.X][end.Y]
		}

		for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newP := Point{p.X + dir.X, p.Y + dir.Y}
			if newP.X < 0 || newP.X >= g.m || newP.Y < 0 || newP.Y >= g.n || g.pointsSet.Has(newP.ToString()) {
				continue
			}
			if !visited[newP.X][newP.Y] {
				queue = append(queue, newP)
				visited[newP.X][newP.Y] = true
				steps[newP.X][newP.Y] = steps[p.X][p.Y] + 1
			}
		}
	}

	return -1
}

func getNewGrid(points []Point) *Grid {
	pointsSet := aoc.Set[string]{}
	for _, p := range points {
		pointsSet.Add(p.ToString())
	}
	return &Grid{points, pointsSet, gridSize, gridSize}
}

func part1(points []Point) int {
	grid := getNewGrid(points[:numBytes])
	grid.Print()
	return grid.GetShortestPath(Point{0, 0}, Point{gridSize - 1, gridSize - 1})
}

func part2(points []Point) {
	// return the first point make BfS fail
	gridPoints := make([]Point, 0)
	for _, p := range points {
		gridPoints = append(gridPoints, p)
		grid := getNewGrid(gridPoints)
		// run bfs
		steps := grid.GetShortestPath(Point{0, 0}, Point{gridSize - 1, gridSize - 1})
		if steps == -1 {
			fmt.Printf("%d,%d\n", p.Y, p.X)
			return
		}
	}

}

func loadInput() []Point {
	lines := aoc.LoadInputLines(inputFileName)
	res := make([]Point, 0)
	for _, line := range lines {
		row := []int{}
		for _, s := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(s)
			row = append(row, n)
		}
		p := Point{row[1], row[0]} // my xy coord is different from the problem's
		res = append(res, p)
	}
	return res
}

func main() {
	input := loadInput()

	fmt.Println(part1(input))
	part2(input)
}

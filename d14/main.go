package main

import (
	"fmt"
	"github.com/swh00tw/aoc"
	"regexp"
	"strconv"
)

var xMax = 101
var yMax = 103

type Robot struct {
	X  int
	Y  int
	Vx int
	Vy int
}

func (r *Robot) Move() {
	//fmt.Println("Move from", r.X, r.Y)
	r.X += r.Vx
	r.Y += r.Vy
	// wrap
	r.X = (r.X + xMax) % xMax
	r.Y = (r.Y + yMax) % yMax
	//fmt.Println("Move to", r.X, r.Y)
}

func loadInput() []*Robot {
	lines := aoc.LoadInputLines("input.txt")
	robots := []*Robot{}
	robotInputRegex := regexp.MustCompile(`p=(-*\d+),(-*\d+)\sv=(-*\d+),(-*\d+)`)
	for _, line := range lines {
		matches := robotInputRegex.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		vx, _ := strconv.Atoi(matches[3])
		vy, _ := strconv.Atoi(matches[4])
		robots = append(robots, &Robot{
			X:  x,
			Y:  y,
			Vx: vx,
			Vy: vy,
		})
	}
	return robots
}

func part1(robots []*Robot) int {
	for i := 0; i < 100; i++ {
		for _, r := range robots {
			r.Move()
		}
	}
	ans := make([]int, 4)
	midX := (xMax - 1) / 2
	midY := (yMax - 1) / 2
	for _, r := range robots {
		x := r.X
		y := r.Y
		if x < midX && y < midY {
			ans[0]++
		} else if x > midX && y > midY {
			ans[1]++
		} else if x < midX && y > midY {
			ans[2]++
		} else if x > midX && y < midY {
			ans[3]++
		}
	}

	return ans[0] * ans[1] * ans[2] * ans[3]
}

func part2(robots []*Robot) {
	n := len(robots)
	for i := 0; i < 10200; i++ {
		graph := make([][]int, yMax)
		for j := 0; j < yMax; j++ {
			graph[j] = make([]int, xMax)
		}

		pointsSet := aoc.Set[string]{}
		for _, r := range robots {
			r.Move()
			graph[r.Y][r.X]++
			key := fmt.Sprintf("%d,%d", r.Y, r.X)
			pointsSet.Add(key)
		}
		// check if some robots overlapped, if yes, skip
		// ref: https://www.reddit.com/r/adventofcode/comments/1hdvhvu/comment/m243der/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
		if pointsSet.Len() < n {
			continue
		}
		fmt.Println("Second: ", i+1, "=========================")

		// print out the graph
		for y := 0; y < yMax; y++ {
			for x := 0; x < xMax; x++ {
				if graph[y][x] == 0 {
					fmt.Printf(".")
				} else {
					fmt.Printf("%d", graph[y][x])
				}
			}
			fmt.Printf("\n")
		}
	}
}

func main() {
	robots := loadInput()
	fmt.Println(part1(robots))

	robots = loadInput()
	part2(robots) // see console to find the answer
}

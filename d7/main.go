package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct {
	value int
	nums  []int
}

func loadInput() []Puzzle {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	puzzles := make([]Puzzle, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		val, _ := strconv.Atoi(line[0])
		rest := strings.Trim(line[1], " ")
		nums := make([]int, 0)
		numbers := strings.Split(rest, " ")
		for _, n := range numbers {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}
		puzzles = append(puzzles, Puzzle{
			value: val,
			nums:  nums,
		})
	}
	return puzzles
}

func solvePuzzle(puzzle Puzzle) bool {
	if len(puzzle.nums) == 0 {
		return false
	}
	target := puzzle.value
	var solve func(curr int, nums []int) bool
	solve = func(curr int, nums []int) bool {
		// base case, when nums is empty
		if len(nums) == 0 {
			return curr == target
		}
		// two case, if one of them true, return true
		rest := nums[1:]
		return solve(curr*nums[0], rest) || solve(curr+nums[0], rest)
	}

	return solve(puzzle.nums[0], puzzle.nums[1:])
}

func part1(puzzles []Puzzle) int {
	ans := 0
	for _, puzzle := range puzzles {
		if solvePuzzle(puzzle) {
			ans += puzzle.value
		}
	}
	return ans
}

func solvePuzzle2(puzzle Puzzle) bool {
	if len(puzzle.nums) == 0 {
		return false
	}
	target := puzzle.value
	var solve func(curr int, nums []int) bool
	solve = func(curr int, nums []int) bool {
		// base case, when nums is empty
		if len(nums) == 0 {
			return curr == target
		}
		// two case, if one of them true, return true
		rest := nums[1:]
		// ADD the third kind of operation: concat
		return solve(curr*nums[0], rest) || solve(curr+nums[0], rest) || solve(concat(curr, nums[0]), rest)
	}

	return solve(puzzle.nums[0], puzzle.nums[1:])
}

func concat(x, y int) int {
	sx := strconv.Itoa(x)
	sy := strconv.Itoa(y)
	n, _ := strconv.Atoi(sx + sy)
	return n
}

func part2(puzzles []Puzzle) int {
	ans := 0
	for _, puzzle := range puzzles {
		if solvePuzzle2(puzzle) {
			ans += puzzle.value
		}
	}
	return ans
}

func main() {
	puzzles := loadInput()

	fmt.Println("part1: ", part1(puzzles))
	fmt.Println("part2: ", part2(puzzles))
}

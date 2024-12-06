package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func loadInput() ([][]int, [][]int) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	scanner := bufio.NewScanner(file)

	hasFinishedSection1 := false
	edges := make([][]int, 0)
	queries := make([][]int, 0)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Trim(text, " ") == "" {
			hasFinishedSection1 = true
			continue
		}
		if hasFinishedSection1 {
			q := []int{}
			nums := strings.Split(text, ",")
			for _, n := range nums {
				num, _ := strconv.Atoi(n)
				q = append(q, num)
			}
			queries = append(queries, q)
		} else {
			e := []int{}
			nums := strings.Split(text, "|")
			for _, n := range nums {
				num, _ := strconv.Atoi(n)
				e = append(e, num)
			}
			edges = append(edges, e)
		}
	}

	return edges, queries
}

func part1(rules [][]int, queries [][]int) int {
	// sol: https://github.com/RemyIsCool/advent-of-code-2024/blob/main/day5/day5.go
	// deep copy
	_q := make([][]int, len(queries))
	for i, q := range queries {
		_q[i] = append([]int{}, append([]int{}, q...)...)
	}
	ans := 0

	for _, q := range _q {
		beforeChange := append([]int{}, q...)
		sort.Slice(q, func(i, j int) bool {
			prev, next := q[i], q[j]
			for _, r := range rules {
				if r[0] == next && r[1] == prev {
					return false
				}
			}
			return true
		})

		if reflect.DeepEqual(beforeChange, q) {
			ans += q[(len(q)-1)/2]
		}
	}

	return ans
}

func part2(rules [][]int, queries [][]int) int {
	// sol: https://github.com/RemyIsCool/advent-of-code-2024/blob/main/day5/day5.go
	// deep copy
	_q := make([][]int, len(queries))
	for i, q := range queries {
		_q[i] = append([]int{}, append([]int{}, q...)...)
	}
	ans := 0

	for _, q := range _q {
		beforeChange := append([]int{}, q...)
		sort.Slice(q, func(i, j int) bool {
			prev, next := q[i], q[j]
			for _, r := range rules {
				if r[0] == next && r[1] == prev {
					return false
				}
			}
			return true
		})

		if !reflect.DeepEqual(beforeChange, q) {
			ans += q[(len(q)-1)/2]
		}
	}

	return ans
}

func main() {
	rules, queries := loadInput()

	fmt.Println(part1(rules, queries))
	fmt.Println(part2(rules, queries))
}

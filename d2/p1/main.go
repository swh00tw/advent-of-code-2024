package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput() ([][]int, error) {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	input := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		nums := make([]int, 0)
		for _, n := range numbers {
			n, _ := strconv.Atoi(n)
			nums = append(nums, n)
		}
		input = append(input, nums)
	}
	return input, nil
}

func validateReport(report []int) bool {
	// validate increasing or decreasing
	for i, n := range report {
		if i == 0 || i == len(report)-1 {
			continue
		}
		if n > report[i-1] && n > report[i+1] {
			return false
		}
		if n < report[i-1] && n < report[i+1] {
			return false
		}
	}

	// sort the slice
	sort.Slice(report, func(i, j int) bool {
		return report[i] < report[j]
	})
	// shold be increasing now
	for i, n := range report {
		if i == 0 {
			continue
		}
		prev := report[i-1]
		if n-prev > 3 {
			return false
		}
		if n-prev < 1 {
			return false
		}
	}
	return true
}

func main() {
	reports, _ := getInput()
	ans := 0
	for _, report := range reports {
		if validateReport(report) {
			ans++
		}
	}

	fmt.Println(ans)
}

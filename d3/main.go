package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func loadInput() []string {
	lines := []string{}
	file, err := os.Open("input.txt")
	if err != nil {
		return []string{}
	}
	defer file.Close()
	// create scanner to read file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func part1(input []string) int {
	ans := 0
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	matches := []string{}
	for _, line := range input {
		m := re.FindAllString(line, -1)
		matches = append(matches, m...)
	}
	for _, match := range matches {
		captureNumRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		numbers := captureNumRegex.FindStringSubmatch(match)
		if len(numbers) == 3 {
			num1, _ := strconv.Atoi(numbers[1])
			num2, _ := strconv.Atoi(numbers[2])
			ans += num1 * num2
		}
	}
	return ans
}

func part2(input []string) int {
	ans := 0
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	matches := []string{}
	for _, line := range input {
		m := re.FindAllString(line, -1)
		matches = append(matches, m...)
	}

	enabled := 1

	for _, match := range matches {
		if match == "do()" {
			enabled = 1
			continue
		}
		if match == "don't()" {
			enabled = 0
			continue
		}
		captureNumRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		numbers := captureNumRegex.FindStringSubmatch(match)
		if len(numbers) == 3 {
			num1, _ := strconv.Atoi(numbers[1])
			num2, _ := strconv.Atoi(numbers[2])
			ans += num1 * num2 * enabled
		}
	}
	return ans
}

func main() {
	lines := loadInput()

	ans1 := part1(lines)
	fmt.Println("Part 1: ", ans1)
	ans2 := part2(lines)
	fmt.Println("Part 2: ", ans2)
}

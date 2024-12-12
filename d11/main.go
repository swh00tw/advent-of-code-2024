package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadInput() []int {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	nums := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		for _, s := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}
	}
	return nums
}

func countDigits(nums int) int {
	str := strconv.Itoa(nums)
	return len(str)
}

func splitStones(num int) (int, int) {
	digits := countDigits(num)
	str := strconv.Itoa(num)
	left := str[:digits/2]
	right := str[digits/2:]
	nl, _ := strconv.Atoi(left)
	nr, _ := strconv.Atoi(right)
	return nl, nr
}

func cached[T comparable, F any](function func(T) F) func(T) F {
	cache := make(map[T]F)
	return func(t T) F {
		if v, ok := cache[t]; ok {
			return v
		}
		ans := function(t)
		cache[t] = ans
		return ans
	}
}

func getNextNums(n int) []int {
	ans := make([]int, 0)
	digits := countDigits(n)
	if n == 0 {
		ans = append(ans, 1)
	} else if digits%2 == 0 {
		l, r := splitStones(n)
		ans = append(ans, l, r)
	} else {
		ans = append(ans, n*2024)
	}
	return ans
}

var cachedGetNextNums = cached(getNextNums)

func blink(freq map[int]int) map[int]int {
	nextFreq := make(map[int]int)
	for k, v := range freq {
		newKeys := cachedGetNextNums(k)
		for _, nk := range newKeys {
			nextFreq[nk] += v
		}
	}
	return nextFreq
}

func part1(nums []int) int {
	freq := map[int]int{}
	for _, n := range nums {
		freq[n]++
	}
	for i := 0; i < 25; i++ {
		freq = blink(freq)
	}
	// count length
	ans := 0
	for _, v := range freq {
		ans += v
	}
	return ans
}

func part2(nums []int) int {
	freq := map[int]int{}
	for _, n := range nums {
		freq[n]++
	}
	for i := 0; i < 75; i++ {
		freq = blink(freq)
	}
	// count length
	ans := 0
	for _, v := range freq {
		ans += v
	}
	return ans
}

func main() {
	nums := loadInput()

	fmt.Println(part1(nums))
	fmt.Println(part2(nums))
}

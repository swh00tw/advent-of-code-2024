package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("p1")
	file, err := os.Open("../p1-input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	ans := 0
	list1 := []int{}
	list2 := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		numbers := strings.Split(text, " ")
		fmt.Println(numbers)
		s1 := numbers[0]
		s2 := numbers[len(numbers)-1]
		n1, _ := strconv.Atoi(s1)
		n2, _ := strconv.Atoi(s2)
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	for i, n1 := range list1 {
		n2 := list2[i]
		ans += abs(n1 - n2)
	}

	fmt.Println("ans: ", ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

package main

import (
	"bufio"
	"fmt"
	"os"
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

	nFreq := make(map[int]int)

	for _, num := range list2 {
		nFreq[num]++
	}
	for _, num := range list1 {
		ans += num * nFreq[num]
	}

	fmt.Println("ans: ", ans)
}

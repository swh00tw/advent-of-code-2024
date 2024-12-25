package main

import (
	"fmt"
	"github.com/swh00tw/aoc"
	"regexp"
	"strconv"
	"strings"
)

var filename = "input.txt"

func loadInput() (map[string]int, [][]string) {
	isFirstPart := true
	lines := aoc.LoadInputLines(filename)

	values := make(map[string]int)
	gates := make([][]string, 0)
	gatesRegex := regexp.MustCompile(`^([a-z\d]{3})\s(AND|OR|XOR)\s([a-z\d]{3})\s->\s([a-z\d]{3})$`)
	for _, line := range lines {
		if line == "" {
			isFirstPart = false
			continue
		}
		if isFirstPart {
			tmp := strings.Split(line, ": ")
			wire, value := tmp[0], tmp[1]
			if value == "1" {
				values[wire] = 1
			} else {
				values[wire] = 0
			}
		} else {
			matches := gatesRegex.FindStringSubmatch(line)
			if len(matches) == 5 {
				gates = append(gates, matches[1:])
			} else {
				fmt.Println("Error parsing line:", line)
			}
		}
	}

	return values, gates
}

func getBinaryNumber(gateVals map[string]int, prefix string) int {
	ans := 0
	bit := 0
	for {
		num := strconv.Itoa(bit)
		if len(num) == 1 {
			num = "0" + num
		}
		key := prefix + num
		if _, ok := gateVals[key]; ok {
			ans = ans | (gateVals[key] << bit)
		} else {
			break
		}
		bit++
	}
	return ans
}

func simulate(_values map[string]int, gates [][]string) int {
	// copy
	values := make(map[string]int)
	for k, v := range _values {
		values[k] = v
	}

	n := len(gates)
	rest := n
	for rest > 0 {
		for _, gate := range gates {
			src1, op, src2, dest := gate[0], gate[1], gate[2], gate[3]
			if _, ok := values[dest]; ok {
				continue
			}
			// fmt.Println(src1, op, src2, dest)
			if _, ok := values[src1]; !ok {
				continue
			}
			if _, ok := values[src2]; !ok {
				continue
			}
			// fmt.Println("Computing", src1, op, src2, dest)
			switch op {
			case "AND":
				values[dest] = values[src1] & values[src2]
			case "OR":
				values[dest] = values[src1] | values[src2]
			case "XOR":
				values[dest] = values[src1] ^ values[src2]
			}
			rest--
		}
	}

	return getBinaryNumber(values, "z")
}

func part1(_values map[string]int, gates [][]string) int {
	return simulate(_values, gates)
}

func main() {
	values, gates := loadInput()

	ans := part1(values, gates)
	fmt.Println("Part1 Answer:", ans)

}

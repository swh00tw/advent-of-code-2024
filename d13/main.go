package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Machine struct {
	a     []int64
	b     []int64
	prize []int64
}

func newMachine() *Machine {
	return &Machine{
		a:     make([]int64, 0),
		b:     make([]int64, 0),
		prize: make([]int64, 0),
	}
}

func (m *Machine) findPossibleAns() (int64, int64, bool) {
	xNum, xDen := m.prize[0]*m.b[1]-m.b[0]*m.prize[1], m.a[0]*m.b[1]-m.b[0]*m.a[1]
	yNum, yDen := m.a[0]*m.prize[1]-m.prize[0]*m.a[1], m.a[0]*m.b[1]-m.b[0]*m.a[1]

	a := xNum / xDen
	b := yNum / yDen
	if a*xDen != xNum || b*yDen != yNum {
		return 0, 0, false
	}
	return a, b, true
}

func part1(machines []Machine) int64 {
	token := int64(0)
	for _, m := range machines {
		a, b, ok := m.findPossibleAns()
		if !ok {
			continue
		}
		token += a*3 + b
	}
	return token
}

func part2(machines []Machine) int64 {
	token := int64(0)
	for _, m := range machines {
		for i, _ := range m.prize {
			m.prize[i] += 10000000000000
		}
		a, b, ok := m.findPossibleAns()
		if !ok {
			continue
		}
		token += a*3 + b
	}
	return token
}

func loadInput() []Machine {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	machines := make([]Machine, 0)
	buttonARegex := regexp.MustCompile(`Button\sA:\sX\+(\d{2}),\sY\+(\d{2})`)
	buttonBRegex := regexp.MustCompile(`Button\sB:\sX\+(\d{2}),\sY\+(\d{2})`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d*), Y=(\d*)`)
	machine := newMachine()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			machines = append(machines, *machine)
			machine = newMachine()
			continue
		}
		if buttonARegex.MatchString(line) {
			matches := buttonARegex.FindStringSubmatch(line)
			a, _ := strconv.Atoi(matches[1])
			b, _ := strconv.Atoi(matches[2])
			machine.a = append(machine.a, int64(a), int64(b))
		} else if buttonBRegex.MatchString(line) {
			matches := buttonBRegex.FindStringSubmatch(line)
			a, _ := strconv.Atoi(matches[1])
			b, _ := strconv.Atoi(matches[2])
			machine.b = append(machine.b, int64(a), int64(b))
		} else if prizeRegex.MatchString(line) {
			matches := prizeRegex.FindStringSubmatch(line)
			a, _ := strconv.Atoi(matches[1])
			b, _ := strconv.Atoi(matches[2])
			machine.prize = append(machine.prize, int64(a), int64(b))
		}
	}
	machines = append(machines, *machine)

	return machines
}

func main() {
	machines := loadInput()

	fmt.Println(part1(machines))
	fmt.Println(part2(machines))
}

// a1 x+b1 y = c1
// a2 x+b2 y = c2
// to get x, c1*b2-c2*b1,

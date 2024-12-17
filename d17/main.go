package main

import (
	"fmt"
	"github.com/swh00tw/aoc"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var inputFileName = "input.txt"

type Computer struct {
	A     int
	B     int
	C     int
	ip    int
	code  []int
	out   []int
	debug bool
}

func (c *Computer) CreateCopy() *Computer {
	return &Computer{
		c.A, c.B, c.C, c.ip, c.code, c.out, c.debug,
	}
}

func (c *Computer) PrintState() {
	fmt.Println("state: ", c.A, c.B, c.C, c.ip)
}

func (c *Computer) GetCode() string {
	// join c.code with ", "
	strs := []string{}
	for _, n := range c.code {
		strs = append(strs, strconv.Itoa(n))
	}
	return strings.Join(strs, ", ")
}

func (c *Computer) GetOutput() string {
	// join c.out with ","
	strs := []string{}
	for _, n := range c.out {
		strs = append(strs, strconv.Itoa(n))
	}
	return strings.Join(strs, ",")
}

func (c *Computer) GetComboOperand(n int) int {
	if n <= 3 {
		return n
	}
	if n == 4 {
		return c.A
	}
	if n == 5 {
		return c.B
	}
	return c.C
}

func (c *Computer) ParseInstruction() {
	opcode := c.code[c.ip]
	comboOperand := c.GetComboOperand(c.code[c.ip+1])
	literalOperand := c.code[c.ip+1]

	switch opcode {
	case 0:
		res := c.A / aoc.IntPow(2, comboOperand)
		c.A = res
		c.ip += 2
	case 1:
		c.B = c.B ^ literalOperand
		c.ip += 2
	case 2:
		c.B = comboOperand % 8
		c.ip += 2
	case 3:
		if c.A == 0 {
			c.ip += 2
		} else {
			c.ip = literalOperand
		}
	case 4:
		c.B = c.B ^ c.C
		c.ip += 2
	case 5:
		c.out = append(c.out, comboOperand%8)
		c.ip += 2
	case 6:
		res := c.A / aoc.IntPow(2, comboOperand)
		c.B = res
		c.ip += 2
	case 7:
		res := c.A / aoc.IntPow(2, comboOperand)
		c.C = res
		c.ip += 2
	}
	if c.debug {
		fmt.Printf("opcode = %d, literal operand = %d, combo operand = %d,\n", opcode, literalOperand, comboOperand)
		c.PrintState()
		fmt.Println()
	}
}

func (c *Computer) Run() {
	for c.ip < len(c.code) {
		c.ParseInstruction()
	}
}

func (c *Computer) Run2() bool {
	for c.ip < len(c.code) {
		c.ParseInstruction()
		n := len(c.out)
		if n > 0 && c.out[n-1] != c.code[n-1] {
			return false
		}
	}
	return true
}

func (c *Computer) Print() {
	for _, n := range c.out {
		fmt.Printf("%d,", n)
	}
}

func part1(com *Computer) {

	//com.debug = true
	com.PrintState()
	fmt.Println()

	com.Run()
	com.Print()
	fmt.Println()
}

func part2(c *Computer) int {
	// some key observation
	/*
			- thinking in bits
			- the length of the output depends on how many bits A have
		    - after each iteration, A /= 8, A got remove 3 bits from the right
			- going reversely, we can search first which A value can have the last digit of program matched
			- then, add 3 bit, try to match last 2 digits
			- then, add 3 bit, try to match last 3 digits
			- and so on
			- until match last n digits (n == len(c.code))
	*/

	searchNums := []int{0}
	for length := 1; length <= len(c.code); length++ {

		nextSearchNums := []int{}
		for _, num := range searchNums {
			// add 3 bits, try all combinations
			for i := 0; i < 8; i++ {
				initA := num*8 + i
				comCopy := c.CreateCopy()
				comCopy.A = initA
				comCopy.Run()
				if slices.Equal(comCopy.out, c.code[len(c.code)-length:]) {
					nextSearchNums = append(nextSearchNums, initA)
				}
			}
		}

		searchNums = nextSearchNums
	}
	return slices.Min(searchNums)
}

func loadInput() *Computer {
	lines := aoc.LoadInputLines(inputFileName)
	firstHalf := true
	registerValueRegex := regexp.MustCompile(`Register\s([A-Z]):\s(\d+)`)
	computer := Computer{
		0, 0, 0, 0, []int{}, []int{}, false,
	}

	for _, line := range lines {
		if line == "" {
			firstHalf = false
			continue
		}
		if firstHalf {
			matches := registerValueRegex.FindStringSubmatch(line)
			if matches[1] == "A" {
				computer.A, _ = strconv.Atoi(matches[2])
			}
			if matches[1] == "B" {
				computer.B, _ = strconv.Atoi(matches[2])
			}
			if matches[1] == "C" {
				computer.C, _ = strconv.Atoi(matches[2])
			}
		} else {
			strs := strings.Split(line, " ")
			numsString := strings.Split(strs[1], ",")
			nums := []int{}
			for _, str := range numsString {
				num, _ := strconv.Atoi(str)
				nums = append(nums, num)
			}
			computer.code = nums
		}
	}
	return &computer
}

func main() {
	computer := loadInput()

	fmt.Println(*computer)
	part1(computer.CreateCopy())
	fmt.Println(part2(computer))
}

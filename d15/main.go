package main

import (
	"fmt"
	"github.com/swh00tw/aoc"
)

// direction enum
type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

var directionsToVec = map[Direction][]int{
	Up:    {-1, 0},
	Right: {0, 1},
	Down:  {1, 0},
	Left:  {0, -1},
}

type Game struct {
	Rx      int // robot x
	Ry      int // robot y
	currMap [][]byte
}

func loadInput(mapper func(byte) []byte) (*Game, []Direction) {
	currMap := [][]byte{}
	moves := []Direction{}
	hasFirstPartFinished := false

	lines := aoc.LoadInputLines("input2.txt")
	for _, line := range lines {
		if line == "" {
			hasFirstPartFinished = true
			continue
		}
		if hasFirstPartFinished {
			for _, b := range []byte(line) {
				switch b {
				case '^':
					moves = append(moves, Up)
				case '>':
					moves = append(moves, Right)
				case 'v':
					moves = append(moves, Down)
				case '<':
					moves = append(moves, Left)
				}
			}
		} else {
			row := []byte{}
			for _, b := range []byte(line) {
				row = append(row, mapper(b)...)
			}
			currMap = append(currMap, row)
		}
	}

	rx := -1
	ry := -1
	m := len(currMap)
	n := len(currMap[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if currMap[i][j] == '@' {
				rx = i
				ry = j
				break
			}
		}
	}

	return &Game{rx, ry, currMap}, moves
}

func (g *Game) MoveRobot(direction Direction) {
	vec := directionsToVec[direction]
	// find "." on the map, if meet box: "O", keep finding, if meet "#", stop
	// if not find, return
	// if find, move all box and robot by one unit
	rx := g.Rx
	ry := g.Ry
	currX := rx + vec[0]
	currY := ry + vec[1]
	for g.currMap[currX][currY] != '.' {
		if g.currMap[currX][currY] == '#' {
			return
		}
		currX = currX + vec[0]
		currY = currY + vec[1]
	}
	// find '.' at currX, currY
	// go reversely to "@", keep swapping
	vX := -vec[0]
	vY := -vec[1]
	for {
		if currX == g.Rx && currY == g.Ry {
			break
		}
		nx := currX + vX
		ny := currY + vY
		g.currMap[nx][ny], g.currMap[currX][currY] = g.currMap[currX][currY], g.currMap[nx][ny]
		currX = nx
		currY = ny
	}
	// update g.Rx, g.Ry
	g.Rx = g.Rx + vec[0]
	g.Ry = g.Ry + vec[1]
}

type Box struct {
	X, Y int
}

func (g *Game) FindBoxes(box Box, dir Direction) ([]Box, bool) {
	//fmt.Println("find boxes for ", box)
	// return boxes, meetWall
	vec := directionsToVec[dir]
	boxes := make([]Box, 0)
	hasWall := false

	nx := box.X + vec[0]
	ny := box.Y + vec[1]
	if dir == Left {
		if g.currMap[nx][ny] == ']' {
			boxes = append(boxes, Box{nx, ny - 1})
		} else if g.currMap[nx][ny] == '#' {
			hasWall = true
		}
	} else if dir == Right {
		if g.currMap[nx][ny+1] == '[' {
			boxes = append(boxes, Box{nx, ny + 2})
		} else if g.currMap[nx][ny+1] == '#' {
			hasWall = true
		}
	} else if dir == Down {
		for k := ny - 1; k <= ny+1; k++ {
			if g.currMap[nx][k] == '[' {
				boxes = append(boxes, Box{nx, k})
			}
			if k >= ny && g.currMap[nx][k] == '#' {
				hasWall = true
			}
		}
	} else if dir == Up {
		for k := ny - 1; k <= ny+1; k++ {
			if g.currMap[nx][k] == '[' {
				boxes = append(boxes, Box{nx, k})
			}
			if k >= ny && g.currMap[nx][k] == '#' {
				hasWall = true
			}
		}
	}
	//fmt.Println("find boxes: ", boxes)
	return boxes, hasWall
}

func (g *Game) MoveRobot2(direction Direction) {
	vec := directionsToVec[direction]
	// find "." on the map, if meet box: "[]", store box in tmp array, next iteration, search box's interface
	// if meet '#', stop
	// move all boxes
	// move robot
	allBoxes := []Box{} // store all boxes
	currBatch := aoc.Set[Box]{}

	cX := g.Rx + vec[0]
	cY := g.Ry + vec[1]
	if g.currMap[cX][cY] == '#' {
		return
	} else if g.currMap[cX][cY] == '.' {
		g.currMap[g.Rx][g.Ry] = '.'
		g.currMap[cX][cY] = '@'
		g.Rx = cX
		g.Ry = cY
		return
	}

	if g.currMap[cX][cY] == '[' {
		currBatch.Add(Box{
			cX, cY,
		})
	} else if g.currMap[cX][cY] == ']' {
		currBatch.Add(Box{
			cX, cY - 1,
		})
	}

	for {
		//fmt.Println("curr batch: ", currBatch)
		nextBatch := aoc.Set[Box]{}
		for _, box := range currBatch.ToArray() {
			boxes, hasWall := g.FindBoxes(box, direction)
			if hasWall {
				return
			}
			for _, b := range boxes {
				nextBatch.Add(b)
			}
		}
		allBoxes = append(allBoxes, currBatch.ToArray()...)
		if nextBatch.Len() == 0 {
			break
		}
		currBatch = nextBatch
	}

	//fmt.Println("Move all boxes: ", allBoxes)
	// move boxes (clear first and redraw)
	for _, box := range allBoxes {
		x, y := box.X, box.Y
		g.currMap[x][y] = '.'
		g.currMap[x][y+1] = '.'
	}
	for _, box := range allBoxes {
		x, y := box.X, box.Y
		nx := x + vec[0]
		ny := y + vec[1]
		g.currMap[nx][ny] = '['
		g.currMap[nx][ny+1] = ']'
	}
	// move robot
	g.currMap[g.Rx][g.Ry] = '.'
	g.Rx = g.Rx + vec[0]
	g.Ry = g.Ry + vec[1]
	g.currMap[g.Rx][g.Ry] = '@'
}

func (g *Game) Print() {
	for _, line := range g.currMap {
		fmt.Println(string(line))
	}
}

func part1(g *Game, moves []Direction) int {
	ans := 0

	for _, d := range moves {
		g.MoveRobot(d)
	}

	g.Print()
	for i, line := range g.currMap {
		for j, cell := range line {
			if cell == 'O' {
				ans += 100*i + j
			}
		}
	}

	return ans
}

func part2(g *Game, moves []Direction) int {
	ans := 0

	fmt.Println("Initial State:")
	g.Print()

	for _, d := range moves {
		fmt.Printf("Moving %v\n", d)
		g.MoveRobot2(d)
		g.Print() // Print after each move to see progression
	}

	fmt.Println("Final State:")
	g.Print()

	for i, line := range g.currMap {
		for j, cell := range line {
			if cell == '[' {
				ans += 100*i + j
			}
		}
	}

	return ans
}

func main() {
	game, moves := loadInput(func(x byte) []byte {
		return []byte{
			x,
		}
	})
	fmt.Println(part1(game, moves))

	game, moves = loadInput(func(x byte) []byte {
		if x == '@' {
			return []byte{
				x,
				'.',
			}
		}
		if x == '#' {
			return []byte{
				'#',
				'#',
			}
		}
		if x == '.' {
			return []byte{
				'.',
				'.',
			}
		}
		return []byte{
			'[',
			']',
		}
	})
	fmt.Println(part2(game, moves)) // didn't work
}

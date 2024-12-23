package main

import (
	"fmt"
	"github.com/swh00tw/aoc"
	"slices"
	"strings"
)

var filename = "input.txt"

func loadInput() [][]string {
	lines := aoc.LoadInputLines(filename)
	connections := make([][]string, len(lines))
	for i, line := range lines {
		connections[i] = strings.Split(line, "-")
	}
	return connections
}

type LAN struct {
	adjList map[string][]string
	nodes   aoc.Set[string]
}

func NewLAN(connections [][]string) *LAN {
	lan := &LAN{
		adjList: make(map[string][]string),
		nodes:   aoc.Set[string]{},
	}

	for _, connection := range connections {
		src, dest := connection[0], connection[1]
		lan.adjList[src] = append(lan.adjList[src], dest)
		lan.adjList[dest] = append(lan.adjList[dest], src)
		lan.nodes.Add(src)
	}

	return lan
}
func (lan *LAN) FindTriplets() []string {
	tripletsToKey := func(triplet []string) string {
		// sort first, then concat
		_triplets := make([]string, len(triplet))
		copy(_triplets, triplet)
		slices.Sort(_triplets)
		return strings.Join(_triplets, ",")
	}

	// looking for sets of three computers where each computer in the set is connected to the other two computers.
	triplets := aoc.Set[string]{}
	for node, neighbors := range lan.adjList {
		for i := 0; i < len(neighbors)-1; i++ {
			for j := i + 1; j < len(neighbors); j++ {
				// if all three computers are connected to each other
				tmp := lan.adjList[neighbors[i]]
				connected := false
				for _, node := range tmp {
					if node == neighbors[j] {
						connected = true
						break
					}
				}
				if !connected {
					continue
				}
				triplet := []string{node, neighbors[i], neighbors[j]}
				key := tripletsToKey(triplet)
				triplets.Add(key)
			}
		}
	}
	return triplets.ToArray()
}
func (lan *LAN) HasEdge(src, dest string) bool {
	for _, neighbor := range lan.adjList[src] {
		if neighbor == dest {
			return true
		}
	}
	return false
}

type UnionFind struct {
	root      []int
	rank      []int
	idxToNode map[int]string
	nodeToIdx map[string]int
	n         int
}

func NewUnionFind(nodes []string) *UnionFind {
	n := len(nodes)
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}
	nodeToIdx := make(map[string]int)
	idxToNode := make(map[int]string)
	for i, node := range nodes {
		nodeToIdx[node] = i
		idxToNode[i] = node
	}
	uf := &UnionFind{
		root:      root,
		rank:      make([]int, n),
		n:         n,
		nodeToIdx: nodeToIdx,
		idxToNode: idxToNode,
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.root[x] != x {
		uf.root[x] = uf.Find(uf.root[x])
	}
	return uf.root[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return
	}
	if uf.rank[rootX] < uf.rank[rootY] {
		rootX, rootY = rootY, rootX
	}
	uf.root[rootY] = rootX
	if uf.rank[rootX] == uf.rank[rootY] {
		uf.rank[rootX]++
	}
}

func (uf *UnionFind) GetSets() map[int][]string {
	sets := make(map[int][]string)
	for i := 0; i < uf.n; i++ {
		root := uf.Find(i)
		sets[root] = append(sets[root], uf.idxToNode[i])
	}
	return sets
}

func part1(lan *LAN) int {
	triplets := lan.FindTriplets()
	sets := [][]string{}
	for _, triplet := range triplets {
		sets = append(sets, strings.Split(triplet, ","))
	}

	cnt := 0
	for _, set := range sets {
		// if any start with 't'
		for _, node := range set {
			if strings.HasPrefix(node, "t") {
				cnt++
				break
			}
		}
	}
	return cnt
}

func part2(lan *LAN) {
	triplets := lan.FindTriplets()
	uf := NewUnionFind(triplets)

	for i := 0; i < len(triplets)-1; i++ {
		for j := i + 1; j < len(triplets); j++ {
			// if two triplets differ by one node && if the new node has connection to all nodes in the two triplets
			nodes := make(map[string]int)
			for _, node := range strings.Split(triplets[i], ",") {
				nodes[node]++
			}
			for _, node := range strings.Split(triplets[j], ",") {
				nodes[node]++
			}
			if len(nodes) == 4 {
				test := []string{}
				for _, node := range strings.Split(triplets[i], ",") {
					delete(nodes, node)
					test = append(test, node)
				}
				for node := range nodes {
					// if connected to all nodes in test
					connected := true
					for _, testNode := range test {
						if !lan.HasEdge(node, testNode) {
							connected = false
							break
						}
					}
					if connected {
						uf.Union(uf.nodeToIdx[triplets[i]], uf.nodeToIdx[triplets[j]])
						break
					}
				}
			}
		}
	}

	sets := uf.GetSets()
	maxSet := []string{}
	for _, set := range sets {
		if len(set) > len(maxSet) {
			maxSet = set
		}
	}
	fmt.Println(maxSet)
	nodes := aoc.Set[string]{}
	for _, t := range maxSet {
		for _, node := range strings.Split(t, ",") {
			nodes.Add(node)
		}
	}
	nodesArr := nodes.ToArray()
	slices.Sort(nodesArr)
	fmt.Println(strings.Join(nodesArr, ","))
}

func main() {
	connections := loadInput()
	lan := NewLAN(connections)

	fmt.Println(part1(lan))
	part2(lan)
}

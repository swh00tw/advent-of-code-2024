package main

import (
	"fmt"
	"github.com/swh00tw/aoc"
	"strconv"
)

var filename = "input.txt"

func loadInput() []int {
	lines := aoc.LoadInputLines(filename)
	secret := []int{}
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		secret = append(secret, n)
	}
	return secret
}

func mix(secret int, subject int) int {
	ans := secret
	ans ^= subject
	return ans
}

func prune(secret int) int {
	ans := secret % 16777216
	return ans
}

func getNextSecret(secret int) int {
	ans := secret
	ans = prune(mix(ans, ans*64))
	ans = prune(mix(ans, ans/32))
	ans = prune(mix(ans, ans*2048))
	return ans
}

func part1(seeds []int) int {
	sum := 0
	for _, seed := range seeds {
		sec := seed
		for i := 0; i < 2000; i++ {
			sec = getNextSecret(sec)
		}
		sum += sec
	}
	return sum
}

func part2(seeds []int) int {
	seqs := [][]int{}
	for _, seed := range seeds {
		seq := []int{seed % 10}
		sec := seed
		for i := 0; i < 2000; i++ {
			sec = getNextSecret(sec)
			seq = append(seq, sec%10)
		}
		seqs = append(seqs, seq)
	}

	queue2Key := func(queue []int) string {
		key := ""
		for _, q := range queue {
			key += strconv.Itoa(q)
			key += ","
		}
		return key
	}
	changeToScore := make(map[string]int)

	for _, seq := range seqs {
		changeToScoreForSeq := make(map[string]int)
		queue := []int{}
		for i := 1; i < len(seq); i++ {
			change := seq[i] - seq[i-1]
			// append to queue
			if len(queue) == 4 {
				// remove first element
				queue = queue[1:]
			}
			queue = append(queue, change)

			// update score, if key conflict, skip
			if len(queue) == 4 {
				key := queue2Key(queue)
				if _, ok := changeToScoreForSeq[key]; ok {
					continue
				} else {
					changeToScoreForSeq[key] = seq[i]
				}
			}
		}
		// write the changeToScoreForSeq to changeToScore
		for k, v := range changeToScoreForSeq {
			changeToScore[k] += v
		}
	}

	maxScore := 0
	maxSeq := ""
	for k, score := range changeToScore {
		if score > maxScore {
			maxScore = score
			maxSeq = k
		}
	}

	fmt.Println(maxSeq)

	return maxScore
}

func main() {
	seeds := loadInput()

	fmt.Println(part1(seeds))
	fmt.Println(part2(seeds))
}

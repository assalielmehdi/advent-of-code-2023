package main

import (
	"assalielmehdi/adventofcode2023/util"
)

func solve1(sc *util.Scanner) any {
	rl := sc.Next()
	next := make(map[string][]string)

	for sc.HasNext() {
		from := sc.Next()
		sc.Next()

		next[from] = []string{sc.Next()[1:4], sc.Next()[:3]}
	}

	steps := 0

	i := 0
	cur := "AAA"
	for {
		if cur == "ZZZ" {
			break
		}

		move := rl[i]
		steps++
		i = (i + 1) % len(rl)

		if move == 'L' {
			cur = next[cur][0]
		} else {
			cur = next[cur][1]
		}
	}

	return steps
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int64) int64 {
	return (a * b) / gcd(a, b)
}

func solve2(sc *util.Scanner) any {
	rl := sc.Next()
	next := make(map[string][]string)
	cur := make([]string, 0)

	for sc.HasNext() {
		from := sc.Next()
		sc.Next()
		next[from] = []string{sc.Next()[1:4], sc.Next()[:3]}

		if from[2] == 'A' {
			cur = append(cur, from)
		}
	}

	answer := int64(1)

	for i := range cur {
		steps := int64(0)

		for j := 0; ; j = (j + 1) % len(rl) {
			steps++

			if rl[j] == 'L' {
				cur[i] = next[cur[i]][0]
			} else {
				cur[i] = next[cur[i]][1]
			}

			if cur[i][2] == 'Z' {
				break
			}
		}

		answer = lcm(answer, steps)
	}

	return answer
}

func main() {
	util.Run("Day 8 | Part 1", solve1, "sample1.txt", "sample12.txt", "input1.txt")
	util.Run("Day 8 | Part 2", solve2, "sample2.txt", "input1.txt")
}

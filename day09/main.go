package main

import (
	"slices"
	"strconv"
	"strings"

	"assalielmehdi/adventofcode2023/util"
)

func solveLine(vals []int) int {
	sum := 0
	lastIdx := len(vals) - 1

	for {
		sum += vals[lastIdx]

		allZero := true

		for i := 0; i < lastIdx; i++ {
			vals[i] = vals[i+1] - vals[i]
			allZero = allZero && vals[i] == 0
		}

		if allZero {
			break
		}

		lastIdx--
	}

	return sum
}

func solve(sc *util.Scanner, reverse bool) any {
	answer := 0

	for sc.HasNextLine() {
		fields := strings.Fields(sc.NextLine())
		vals := make([]int, 0, len(fields))

		for _, field := range fields {
			val, _ := strconv.Atoi(field)
			vals = append(vals, val)
		}

		if reverse {
			slices.Reverse(vals)
		}
		answer += solveLine(vals)
	}

	return answer
}

func main() {
	util.RunAll("Day 9 - Part 1", func(sc *util.Scanner) any { return solve(sc, false) })
	util.RunAll("Day 9 - Part 2", func(sc *util.Scanner) any { return solve(sc, true) })
}

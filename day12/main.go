package main

import (
	"assalielmehdi/adventofcode2023/util"
	"strconv"
	"strings"
)

func toInts(fields []string) []int {
	ints := make([]int, 0, len(fields))
	for _, field := range fields {
		val, _ := strconv.Atoi(field)
		ints = append(ints, val)
	}
	return ints
}

func duplicate(springs string, damaged []int, n int) (string, []int) {
	dupSprings := ""
	dupDamaged := make([]int, 0, n*len(damaged))

	for i := 0; i < n; i++ {
		if i > 0 {
			dupSprings += "?"
		}
		dupSprings += springs
		dupDamaged = append(dupDamaged, damaged...)
	}

	return dupSprings, dupDamaged
}

func dp(i int, springs string, j int, damaged []int, memo [][]int) int {
	if j >= len(damaged) {
		allDots := true

		for k := i; allDots && k < len(springs); k++ {
			allDots = springs[k] == '.' || springs[k] == '?'
		}

		if allDots {
			return 1
		}

		return 0
	}

	if i >= len(springs) {
		return 0
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	if springs[i] == '.' {
		memo[i][j] = dp(i+1, springs, j, damaged, memo)
		return memo[i][j]
	}

	count := 0

	if springs[i] == '?' {
		count += dp(i+1, springs, j, damaged, memo)
	}

	canSub := 0
	for k := 0; k < damaged[j] && i+k < len(springs); k++ {
		if springs[i+k] == '?' || springs[i+k] == '#' {
			canSub++
		}
	}

	if canSub != damaged[j] {
		memo[i][j] = count
		return memo[i][j]
	}

	if i+damaged[j] >= len(springs) || springs[i+damaged[j]] == '.' || springs[i+damaged[j]] == '?' {
		count += dp(i+damaged[j]+1, springs, j+1, damaged, memo)
	}

	memo[i][j] = count
	return memo[i][j]
}

func solve(sc *util.Scanner) any {
	sum := 0

	for sc.HasNextLine() {
		line := strings.Fields(sc.NextLine())

		springs := line[0]
		damaged := toInts(strings.Split(line[1], ","))

		springs, damaged = duplicate(springs, damaged, 5)

		memo := make([][]int, 0, len(springs))
		for i := 0; i < len(springs); i++ {
			memo = append(memo, make([]int, len(damaged)))
			for j := 0; j < len(memo[i]); j++ {
				memo[i][j] = -1
			}
		}

		sum += dp(0, springs, 0, damaged, memo)
	}

	return sum
}

func main() {
	util.RunAll("Day 12", solve)
}

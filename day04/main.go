package main

import (
	"strconv"

	"assalielmehdi/adventofcode2023/util"
)

func parseCard(str []byte) (*util.Set[int], *util.Set[int]) {
	sc := util.NewScanner(str)

	sc.Next() // Card
	sc.Next() // XX:

	winning, guess := util.NewSet[int](), util.NewSet[int]()
	current := winning

	for sc.HasNext() {
		token := sc.Next()

		val, err := strconv.Atoi(string(token))
		if err == nil {
			current.Add(val)
			continue
		}

		current = guess
	}

	return winning, guess
}

func scratchCard(str []byte) int {
	winning, guess := parseCard(str)

	guess.Retain(winning)

	return guess.Size()
}

func solve1(sc *util.Scanner) any {
	sum := 0

	for sc.HasNextLine() {
		count := scratchCard(sc.NextLineBytes())

		if count >= 1 {
			sum += 1 << (count - 1)
		}
	}

	return sum
}

func dp(cards []int, i int, memo map[int]int) int {
	if memo[i] != 0 {
		return memo[i]
	}

	memo[i] = 1

	for j := 0; i+j+1 < len(cards) && j < cards[i]; j++ {
		memo[i] += dp(cards, i+j+1, memo)
	}

	return memo[i]
}

func solve2(sc *util.Scanner) any {
	cards := make([]int, 0)
	memo := make(map[int]int)

	for sc.HasNextLine() {
		cards = append(cards, scratchCard(sc.NextLineBytes()))
	}

	sum := 0
	for i := range cards {
		sum += dp(cards, i, memo)
	}

	return sum
}

func main() {
	util.RunAll("Day 4 - Part 1", solve1)
	util.RunAll("Day 4 - Part 2", solve2)
}

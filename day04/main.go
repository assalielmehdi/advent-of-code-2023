package main

import (
	util_ds "assalielmehdi/adventofcode2023/pkg/ds"
	util_fp "assalielmehdi/adventofcode2023/pkg/fp"
	util_io "assalielmehdi/adventofcode2023/pkg/io"
	util_map "assalielmehdi/adventofcode2023/pkg/map"
	"fmt"
	"strconv"
)

func parseCard(str []byte) (*util_ds.Set[int], *util_ds.Set[int]) {
	tokzr := util_io.NewTokenizer(str)

	tokzr.Next() // Card
	tokzr.Next() // XX:

	winning, guess := util_ds.NewSet[int](), util_ds.NewSet[int]()
	current := winning

	for tokzr.HasNext() {
		token := tokzr.Next()

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

func solve1() {
	lines := util_io.ReadLines("input1.txt")

	sum := 0

	for _, line := range lines {
		count := scratchCard(line)

		if count >= 1 {
			sum += 1 << (count - 1)
		}
	}

	fmt.Println(sum)
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

func solve2() {
	lines := util_io.ReadLines("input1.txt")
	cards := util_fp.Map(lines, func(line []byte) int { return scratchCard(line) })
	memo := make(map[int]int)

	util_fp.ForEach(cards, func(i, card int) { dp(cards, i, memo) })

	fmt.Println(util_map.Sum(memo))
}

func main() {
	solve1()
	solve2()
}

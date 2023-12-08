package main

import (
	"assalielmehdi/adventofcode2023/util"
	"strconv"
	"unicode"
)

func parseInt(str string, i int) (int, int) {
	substr := make([]byte, 0)

	for unicode.IsDigit(rune(str[i])) {
		substr = append(substr, str[i])
		i++
	}

	v, _ := strconv.Atoi(string(substr))

	return i, v
}

func parseColor(str string, i int) (int, int, int) {
	i, count := parseInt(str, i)
	i++

	if str[i] == 'r' {
		return i + 3, 0, count
	}

	if str[i] == 'g' {
		return i + 5, 1, count
	}

	return i + 4, 2, count
}

func parseRound(str string, i int) (int, []int) {
	round := make([]int, 3)

	for {
		j, color, count := parseColor(str, i)
		i = j

		round[color] = count

		if i == len(str) {
			break
		}

		if str[i] == ';' {
			i += 2
			break
		}

		i += 2
	}

	return i, round
}

func parseGame(str string) [][]int {
	i := 0
	for str[i] != ':' {
		i++
	}

	i += 2

	rounds := make([][]int, 0)

	for {
		j, round := parseRound(str, i)
		i = j

		rounds = append(rounds, round)

		if i == len(str) {
			break
		}
	}

	return rounds
}

func isGamePossible(maxRed, maxGreen, maxBlue int, game [][]int) bool {
	ok := true

	for _, round := range game {
		ok = ok && round[0] <= maxRed && round[1] <= maxGreen && round[2] <= maxBlue

		if !ok {
			break
		}
	}

	return ok
}

func minPower(game [][]int) int {
	maxRed, maxGreen, maxBlue := 0, 0, 0

	for _, round := range game {
		maxRed = max(maxRed, round[0])
		maxGreen = max(maxGreen, round[1])
		maxBlue = max(maxBlue, round[2])
	}

	return maxRed * maxGreen * maxBlue
}

func solve1(sc *util.Scanner) any {
	sum := 0
	i := 1

	for sc.HasNextLine() {
		game := parseGame(sc.NextLine())

		if isGamePossible(12, 13, 14, game) {
			sum += i
		}

		i++
	}

	return sum
}

func solve2(sc *util.Scanner) any {
	sum := 0

	for sc.HasNextLine() {
		game := parseGame(sc.NextLine())

		sum += minPower(game)
	}

	return sum
}

func main() {
	util.RunAll("Day 2 - Part 1", solve1)
	util.RunAll("Day 2 - Part 2", solve2)
}

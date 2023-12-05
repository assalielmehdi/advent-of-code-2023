package main

import (
	"fmt"
	"strconv"
	"unicode"

	"assalielmehdi/adventofcode2023/util"
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

func main() {
	lines := util.NewFileIterator("sample1.txt")
	defer lines.Close()

	sum := 0

	for lines.HasNext() {
		line := lines.Next()
		game := parseGame(line)

		sum += minPower(game)
	}

	fmt.Println(sum)
}

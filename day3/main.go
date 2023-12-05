package main

import (
	util_io "assalielmehdi/adventofcode2023/pkg/io"
	"fmt"
	"strconv"
	"unicode"
)

func isNextTo(grid [][]byte, i, j int, predicate func(byte) bool) bool {
	next := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for _, nextIJ := range next {
		nextI, nextJ := i+nextIJ[0], j+nextIJ[1]

		if nextI >= 0 && nextI < len(grid) && nextJ >= 0 && nextJ < len(grid[0]) && predicate(grid[nextI][nextJ]) {
			return true
		}
	}

	return false
}

func parseNumber(grid [][]byte, i, j int) int {
	if !unicode.IsDigit(rune(grid[i][j])) {
		return 0
	}

	isNextToSymbol := false
	valueStr := make([]byte, 0)

	isSymbol := func(ch byte) bool {
		return !unicode.IsDigit(rune(ch)) && ch != '.'
	}

	for j < len(grid[0]) && unicode.IsDigit(rune(grid[i][j])) {
		isNextToSymbol = isNextToSymbol || isNextTo(grid, i, j, isSymbol)

		valueStr = append(valueStr, grid[i][j])
		grid[i][j] = '.'
		j++
	}

	if !isNextToSymbol {
		return 0
	}

	value, _ := strconv.Atoi(string(valueStr))

	return value
}

func solve1() {
	grid := make([][]byte, 0)
	lines := util_io.ReadLines("input1.txt")

	for _, line := range lines {
		grid = append(grid, line)
	}

	sum := int64(0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			sum += int64(parseNumber(grid, i, j))
		}
	}

	fmt.Println(sum)
}

func parseNearGears(grid [][]byte, i, j int) [][]int {
	next := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	gearsPositions := make([][]int, 0)

	for _, nextIJ := range next {
		nextI, nextJ := i+nextIJ[0], j+nextIJ[1]

		if nextI >= 0 && nextI < len(grid) && nextJ >= 0 && nextJ < len(grid[0]) && grid[nextI][nextJ] == '*' {
			gearsPositions = append(gearsPositions, []int{nextI, nextJ})
		}
	}

	return gearsPositions
}

func parseGear(grid [][]byte, i, j int, ratioMemo map[string][]int64) {
	if !unicode.IsDigit(rune(grid[i][j])) {
		return
	}

	nearGears := make(map[string]bool, 0)
	valueStr := make([]byte, 0)

	for j < len(grid[0]) && unicode.IsDigit(rune(grid[i][j])) {
		nearGearsPositions := parseNearGears(grid, i, j)

		for _, position := range nearGearsPositions {
			key := fmt.Sprintf("%d,%d", position[0], position[1])
			nearGears[key] = true
		}

		valueStr = append(valueStr, grid[i][j])
		grid[i][j] = '.'
		j++
	}

	value, _ := strconv.Atoi(string(valueStr))

	for key := range nearGears {
		ratios, exists := ratioMemo[key]
		if !exists {
			ratios = make([]int64, 0)
		}

		ratioMemo[key] = append(ratios, int64(value))
	}
}

func solve2() {
	grid := make([][]byte, 0)
	lines := util_io.ReadLines("input1.txt")

	for _, line := range lines {
		grid = append(grid, line)
	}

	ratioMemo := make(map[string][]int64)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			parseGear(grid, i, j, ratioMemo)
		}
	}

	sum := int64(0)

	for _, ratios := range ratioMemo {
		if len(ratios) == 2 {
			product := int64(1)

			for _, ratio := range ratios {
				product *= ratio
			}

			sum += product
		}
	}

	fmt.Println(sum)
}

func main() {
	solve1()
	solve2()
}

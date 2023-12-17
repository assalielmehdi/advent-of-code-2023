package main

import (
	"assalielmehdi/adventofcode2023/util"
	"slices"
)

var dirs = map[byte][][][]int{
	'.':  {{{1, 0, 0}}, {{0, -1, 1}}, {{-1, 0, 2}}, {{0, 1, 3}}},
	'-':  {{{0, 1, 3}, {0, -1, 1}}, {{0, -1, 1}}, {{0, 1, 3}, {0, -1, 1}}, {{0, 1, 3}}},
	'\\': {{{0, 1, 3}}, {{-1, 0, 2}}, {{0, -1, 1}}, {{1, 0, 0}}},
	'/':  {{{0, -1, 1}}, {{1, 0, 0}}, {{0, 1, 3}}, {{-1, 0, 2}}},
	'|':  {{{1, 0, 0}}, {{-1, 0, 2}, {1, 0, 0}}, {{-1, 0, 2}}, {{-1, 0, 2}, {1, 0, 0}}},
}

func next(grid [][]byte, i, j, from int) [][]int {
	nextij := make([][]int, 0)
	for _, dij := range dirs[grid[i][j]][from] {
		nextij = append(nextij, []int{i + dij[0], j + dij[1], dij[2]})
	}
	return nextij
}

func dfs(grid [][]byte, i, j, from int, vis [][][]bool) int {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || vis[i][j][from] {
		return 0
	}
	count := 0
	if !slices.Contains(vis[i][j], true) {
		count++
	}
	vis[i][j][from] = true
	for _, nextij := range next(grid, i, j, from) {
		count += dfs(grid, nextij[0], nextij[1], nextij[2], vis)
	}
	return count
}

func reset(vis [][][]bool) [][][]bool {
	for _, line := range vis {
		for _, col := range line {
			for k := 0; k < 4; k++ {
				col[k] = false
			}
		}
	}
	return vis
}

func solve(sc *util.Scanner) any {
	grid := make([][]byte, 0)
	vis := make([][][]bool, 0)
	for sc.HasNextLine() {
		line := sc.NextLineBytes()
		grid = append(grid, line)
		visLine := make([][]bool, 0, len(line))
		for i := 0; i < len(line); i++ {
			visLine = append(visLine, make([]bool, 4))
		}
		vis = append(vis, visLine)
	}
	answer := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if i == 0 {
				answer = max(answer, dfs(grid, i, j, 0, reset(vis)))
			}
			if i == len(grid)-1 {
				answer = max(answer, dfs(grid, i, j, 2, reset(vis)))
			}
			if j == 0 {
				answer = max(answer, dfs(grid, i, j, 3, reset(vis)))
			}
			if j == len(grid[i])-1 {
				answer = max(answer, dfs(grid, i, j, 1, reset(vis)))
			}
		}
	}
	return answer
}

func main() {
	util.RunAll("Day 16", solve)
}

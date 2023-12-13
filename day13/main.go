package main

import (
	"slices"

	"assalielmehdi/adventofcode2023/util"
)

func findMirrors(grid [][]byte) [][]int {
	mirrors := [][]int{}

	for i := 0; i < len(grid)-1; i++ {
		iok := true

		for j := 0; iok && j < len(grid[i]); j++ {
			jok := true

			il, ir := i, i+1
			for jok && il >= 0 && ir < len(grid) {
				jok = grid[il][j] == grid[ir][j]

				il--
				ir++
			}

			iok = jok
		}

		if iok {
			mirrors = append(mirrors, []int{100 * (i + 1), 0})
		}
	}

	for j := 0; j < len(grid[0])-1; j++ {
		jok := true

		for i := 0; jok && i < len(grid); i++ {
			iok := true

			jl, jr := j, j+1
			for iok && jl >= 0 && jr < len(grid[0]) {
				iok = grid[i][jl] == grid[i][jr]

				jl--
				jr++
			}

			jok = iok
		}

		if jok {
			mirrors = append(mirrors, []int{j + 1, 1})
		}
	}

	return mirrors
}

func findSmudge(grid [][]byte) int {
	oldMirrors := findMirrors(grid)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			oldCh := grid[i][j]

			if oldCh == '.' {
				grid[i][j] = '#'
			} else {
				grid[i][j] = '.'
			}

			newMirrors := findMirrors(grid)
			grid[i][j] = oldCh

			for _, newMirror := range newMirrors {
				if !slices.ContainsFunc(oldMirrors, func(oldMirror []int) bool { return slices.Equal(oldMirror, newMirror) }) {
					return newMirror[0]
				}
			}
		}
	}

	return 0
}

func solve(sc *util.Scanner) any {
	grid := make([][]byte, 0)
	answer := 0

	for sc.HasNextLine() {
		line := sc.NextLineBytes()
		if len(line) == 0 {
			answer += findSmudge(grid)

			grid = make([][]byte, 0)
			continue
		}

		grid = append(grid, line)
	}

	answer += findSmudge(grid)

	return answer
}

func main() {
	util.RunAll("Day 13", solve)
}

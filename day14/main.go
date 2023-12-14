package main

import (
	"slices"

	"assalielmehdi/adventofcode2023/util"
)

func rotate(grid [][]byte) [][]byte {
	rot := make([][]byte, 0, len(grid[0]))
	for j := 0; j < len(grid[0]); j++ {
		rot = append(rot, make([]byte, len(grid)))
		for i := 0; i < len(grid); i++ {
			rot[j][i] = grid[i][j]
		}
		slices.Reverse(rot[j])
	}
	return rot
}

func tiltNorth(grid [][]byte) {
	for j := 0; j < len(grid[0]); j++ {
		imin := -1
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == '.' {
				if imin == -1 {
					imin = i
				}
			} else if grid[i][j] == '#' {
				imin = -1
			} else {
				if imin != -1 {
					grid[imin][j] = grid[i][j]
					grid[i][j] = '.'
					imin++
				}
			}
		}
	}
}

func load(grid [][]byte) int {
	answer := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'O' {
				answer += len(grid) - i
			}
		}
	}
	return answer
}

func cycle(grid [][]byte) [][]byte {
	for i := 0; i < 4; i++ {
		tiltNorth(grid)
		grid = rotate(grid)
	}
	return grid
}

func findMod(vals []int) int {
	for mod := 1; mod <= len(vals)/2; mod++ {
		can := true
		for i := 0; can && i < mod; i++ {
			can = vals[i] == vals[i+mod]
		}
		if !can {
			continue
		}
		ok := true
		for i := 0; ok && i < mod; i++ {
			for j := 1; ok && j < len(vals)/mod; j++ {
				if i+j*mod < len(vals) {
					ok = vals[i] == vals[i+j*mod]
				}
			}
		}
		if ok {
			return mod
		}
	}
	return -1
}

func solve(sc *util.Scanner) any {
	grid := make([][]byte, 0)
	for sc.HasNextLine() {
		grid = append(grid, sc.NextLineBytes())
	}

	// Run arbitrary # of cycles
	loads := make([]int, 0, int(1e3))
	for i := 0; i < int(1e3); i++ {
		grid = cycle(grid)
		loads = append(loads, load(grid))
	}

	imod, mod := -1, -1
	for i := 0; i < len(loads); i++ {
		submod := findMod(loads[i:])
		if submod > mod {
			imod = i
			mod = submod
		}
	}

	ianswer := (int(1e9) - (imod + 1)) % mod

	return loads[imod+ianswer]
}

func main() {
	util.RunAll("Day 14", solve)
}

package main

import (
	"assalielmehdi/adventofcode2023/util"
	"slices"
)

func findCycle(grid [][]byte, si, sj int) [][]int {
	cycle := [][]int{{si, sj}}
	var dij []int

	if si-1 >= 0 && (grid[si-1][sj] == '|' || grid[si-1][sj] == 'F' || grid[si-1][sj] == '7') {
		dij = []int{-1, 0}
	} else if sj+1 < len(grid[0]) && (grid[si][sj+1] == '-' || grid[si][sj+1] == 'J' || grid[si][sj+1] == '7') {
		dij = []int{0, 1}
	} else if si+1 < len(grid) && (grid[si+1][sj] == '|' || grid[si+1][sj] == 'L' || grid[si+1][sj] == 'J') {
		dij = []int{1, 0}
	} else if sj-1 >= 0 && (grid[si][sj-1] == '-' || grid[si][sj+1] == 'F' || grid[si][sj+1] == 'L') {
		dij = []int{0, -1}
	}

	last := cycle[0]
	for {
		cur := []int{last[0] + dij[0], last[1] + dij[1]}
		if cur[0] == si && cur[1] == sj {
			break
		}

		cycle = append(cycle, cur)
		last = cur

		if dij[0] == 1 && dij[1] == 0 {
			if grid[cur[0]][cur[1]] == '|' {
				dij = []int{1, 0}
			} else if grid[cur[0]][cur[1]] == 'L' {
				dij = []int{0, 1}
			} else if grid[cur[0]][cur[1]] == 'J' {
				dij = []int{0, -1}
			}
		} else if dij[0] == 0 && dij[1] == -1 {
			if grid[cur[0]][cur[1]] == '-' {
				dij = []int{0, -1}
			} else if grid[cur[0]][cur[1]] == 'L' {
				dij = []int{-1, 0}
			} else if grid[cur[0]][cur[1]] == 'F' {
				dij = []int{1, 0}
			}
		} else if dij[0] == -1 && dij[1] == 0 {
			if grid[cur[0]][cur[1]] == '|' {
				dij = []int{-1, 0}
			} else if grid[cur[0]][cur[1]] == 'F' {
				dij = []int{0, 1}
			} else if grid[cur[0]][cur[1]] == '7' {
				dij = []int{0, -1}
			}
		} else if dij[0] == 0 && dij[1] == 1 {
			if grid[cur[0]][cur[1]] == '-' {
				dij = []int{0, 1}
			} else if grid[cur[0]][cur[1]] == 'J' {
				dij = []int{-1, 0}
			} else if grid[cur[0]][cur[1]] == '7' {
				dij = []int{1, 0}
			}
		}
	}

	return cycle
}

func solve1(sc *util.Scanner) any {
	grid := make([][]byte, 0)
	si, sj := -1, -1

	for i := 0; sc.HasNextLine(); i++ {
		grid = append(grid, sc.NextLineBytes())

		if si == -1 {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == 'S' {
					si, sj = i, j
				}
			}
		}
	}

	cycle := findCycle(grid, si, sj)

	return (len(cycle) + 1) / 2
}

func getArea(pts [][]int, n int) int {
	sum := 0
	for i := 0; i < len(pts); i++ {
		pt1, pt2 := pts[i], pts[(i+1)%len(pts)]
		x1, y1 := pt1[1], n-1-pt1[0]
		x2, y2 := pt2[1], n-1-pt2[0]

		sum += x1*y2 - x2*y1
	}
	return sum / 2
}

func solve2(sc *util.Scanner) any {
	grid := make([][]byte, 0)
	si, sj := -1, -1

	for i := 0; sc.HasNextLine(); i++ {
		grid = append(grid, sc.NextLineBytes())

		if si == -1 {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == 'S' {
					si, sj = i, j
				}
			}
		}
	}

	cycle := findCycle(grid, si, sj)
	area := getArea(cycle, len(grid))
	slices.Reverse(cycle)
	area = max(area, getArea(cycle, len(grid)))

	answer := area + 1 - len(cycle)/2

	return answer
}

func main() {
	util.Run("Day 10 - Part 1", solve1, "sample11.txt", "sample12.txt", "input1.txt")
	util.Run("Day 10 - Part 2", solve2, "sample20.txt", "sample21.txt", "sample22.txt", "sample23.txt", "sample24.txt", "input1.txt")
}

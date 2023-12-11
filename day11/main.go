package main

import (
	"assalielmehdi/adventofcode2023/util"
)

func bfs(ri, rj int, grid [][]byte, sumRow, sumCol []int, expVal int) int {
	vis := make([][]bool, 0, len(grid))
	for i := 0; i < len(grid); i++ {
		vis = append(vis, make([]bool, len(grid[0])))
	}

	answer := 0
	q := [][]int{{ri, rj, 0}}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		i, j, dist := cur[0], cur[1], cur[2]

		if vis[i][j] {
			continue
		}
		vis[i][j] = true

		if grid[i][j] != '.' && i*len(grid)+j > ri*len(grid)+rj {
			answer += dist
		}

		di := []int{-1, 0, 1, 0}
		dj := []int{0, 1, 0, -1}

		for k := 0; k < 4; k++ {
			ii := i + di[k]
			jj := j + dj[k]
			nextDist := dist + 1

			if ii != i && sumRow[i] == 0 {
				nextDist += expVal
			}

			if jj != j && sumCol[j] == 0 {
				nextDist += expVal
			}

			if ii >= 0 && ii < len(grid) && jj >= 0 && jj < len(grid[0]) && !vis[ii][jj] {
				q = append(q, []int{ii, jj, nextDist})
			}
		}
	}

	return answer
}

func solve(sc *util.Scanner) any {
	grid := make([][]byte, 0)
	for sc.HasNextLine() {
		grid = append(grid, sc.NextLineBytes())
	}

	pos := make([][]int, 0)
	sumRow := make([]int, len(grid))
	sumCol := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '.' {
				sumRow[i]++
				sumCol[j]++
				pos = append(pos, []int{i, j})
			}
		}
	}

	answer := 0
	for _, ij := range pos {
		answer += bfs(ij[0], ij[1], grid, sumRow, sumCol, 1000000-1)
	}

	return answer
}

func main() {
	util.RunAll("Day 11", solve)
}

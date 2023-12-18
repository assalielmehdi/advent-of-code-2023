package main

import (
	"assalielmehdi/adventofcode2023/util"
	"container/heap"
)

const Inf = int(1e9)

type state struct {
	i, j, from, strike, dist int
}

func (s1 *state) Less(s2 *state) bool {
	return s1.dist < s2.dist
}

func next(i, j, from, strike int) [][4]int {
	dir := [][][]int{
		{{1, 0, 0}, {0, -1, 1}, {0, 1, 3}},
		{{0, -1, 1}, {-1, 0, 2}, {1, 0, 0}},
		{{-1, 0, 2}, {0, 1, 3}, {0, -1, 1}},
		{{0, 1, 3}, {1, 0, 0}, {-1, 0, 2}},
	}
	nextDir := dir[from]
	if strike < 4 {
		return [][4]int{{i + nextDir[0][0], j + nextDir[0][1], nextDir[0][2], strike + 1}}
	}
	if strike == 10 {
		nextDir = nextDir[1:]
	}
	nextPos := make([][4]int, 0, len(nextDir))
	for k, ijf := range nextDir {
		pos := [4]int{i + ijf[0], j + ijf[1], ijf[2], 1}
		if strike < 10 && k == 0 {
			pos[3] = strike + 1
		}
		nextPos = append(nextPos, pos)
	}
	return nextPos
}

func newDist() [4][11]int {
	dist := [4][11]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 11; j++ {
			dist[i][j] = Inf
		}
	}
	return dist
}

func solve1(sc *util.Scanner) any {
	grid := make([][]int, 0)
	dist := make([][][4][11]int, 0)
	for i := 0; sc.HasNextLine(); i++ {
		line := sc.NextLine()
		grid = append(grid, make([]int, 0, len(line)))
		dist = append(dist, make([][4][11]int, 0, len(line)))
		for _, d := range line {
			grid[i] = append(grid[i], int(d-'0'))
			dist[i] = append(dist[i], newDist())
		}
	}
	n, m := len(grid), len(grid[0])
	pq := util.NewPriorityQueue[*state]()
	heap.Init(pq)
	heap.Push(pq, &state{0, 0, 0, 0, 0})
	heap.Push(pq, &state{0, 0, 3, 0, 0})
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*state)
		if cur.i == -1 || cur.j == -1 || cur.i == n || cur.j == m {
			continue
		}
		if dist[cur.i][cur.j][cur.from][cur.strike] != Inf {
			continue
		}
		dist[cur.i][cur.j][cur.from][cur.strike] = cur.dist
		for _, next := range next(cur.i, cur.j, cur.from, cur.strike) {
			if next[0] == -1 || next[1] == -1 || next[0] == n || next[1] == m {
				continue
			}
			if cur.dist+grid[next[0]][next[1]] < dist[next[0]][next[1]][next[2]][next[3]] {
				heap.Push(pq, &state{next[0], next[1], next[2], next[3], cur.dist + grid[next[0]][next[1]]})
			}
		}
	}
	answer := Inf
	for i := 0; i < 4; i++ {
		for j := 0; j < 11; j++ {
			if j >= 4 {
				answer = min(answer, dist[n-1][m-1][i][j])
			}
		}
	}
	return answer
}

func main() {
	util.Run("Day 17 - Part 2", solve1, "input1.txt")
}

package main

import (
	"slices"
	"strconv"
	"strings"

	"assalielmehdi/adventofcode2023/util"
)

type inst struct {
	dir util.Dir
	len int64
}

func shoelace(pts []*util.Pair[int64, int64]) int64 {
	a := int64(0)
	for i := 0; i < len(pts); i++ {
		pt1, pt2 := pts[i], pts[(i+1)%len(pts)]
		a += pt1.First*pt2.Second - pt1.Second*pt2.First
	}
	return a / 2
}

func picks(a, b int64) int64 {
	return (a - b/2) + 1
}

func solve(insts []*inst) int64 {
	b := int64(0)
	pts := make([]*util.Pair[int64, int64], 0)
	cur := util.Pair[int64, int64]{First: 0, Second: 0}
	for _, inst := range insts {
		b += inst.len
		next := util.Pair[int64, int64]{First: cur.First, Second: cur.Second}
		switch inst.dir {
		case util.DirUp:
			next.First -= inst.len
		case util.DirRight:
			next.Second += inst.len
		case util.DirDown:
			next.First += inst.len
		case util.DirLeft:
			next.Second -= inst.len
		}
		pts = append(pts, &next)
		cur = next
	}
	a := -shoelace(pts)
	return picks(a, b) + b
}

func solve2(sc *util.Scanner) any {
	insts := make([]*inst, 0)
	for sc.HasNextLine() {
		line := sc.NextLineBytes()
		inst := inst{
			dir: util.Dir(line[len(line)-1] - '0'),
		}
		len, _ := strconv.ParseInt(string(line[:len(line)-1]), 16, 64)
		inst.len = len
		insts = append(insts, &inst)
	}
	return solve(insts)
}

func solve1(sc *util.Scanner) any {
	insts := make([]*inst, 0)
	for sc.HasNextLine() {
		fields := strings.Fields(sc.NextLine())
		inst := inst{
			dir: util.Dir(slices.Index([]string{"R", "D", "L", "U"}, fields[0])),
		}
		len, _ := strconv.ParseInt(fields[1], 10, 64)
		inst.len = len
		insts = append(insts, &inst)
	}
	return solve(insts)
}

func main() {
	util.Run("Day 18 - Part 1", solve1, "sample1.txt", "input1.txt")
	util.Run("Day 18 - Part 2", solve2, "sample2.txt", "input2.txt")
}

package main

import (
	"assalielmehdi/adventofcode2023/util"
	"slices"
	"strconv"
	"strings"
)

func toInts(fields []string) []int {
	ints := make([]int, 0, len(fields))
	for _, field := range fields {
		val, _ := strconv.Atoi(field)
		ints = append(ints, val)
	}
	return ints
}

func applyMapp(src [][]int, mapp [][]int) [][]int {
	dst := make([][]int, 0)
	isrc, imapp := 0, 0

	for isrc < len(src) && imapp < len(mapp) {
		mappInter := mapp[imapp]
		srcInter := []int{src[isrc][0] + src[isrc][2], src[isrc][1] + src[isrc][2]}

		if mappInter[0] > srcInter[1] {
			dst = append(dst, src[isrc])
			isrc++
			continue
		}

		if mappInter[1] < srcInter[0] {
			imapp++
			continue
		}

		leftInter := []int{srcInter[0], max(mappInter[0], srcInter[0]) - 1}
		midInter := []int{leftInter[1] + 1, min(mappInter[1], srcInter[1])}
		rightInter := []int{midInter[1] + 1, srcInter[1]}

		if leftInter[0] <= leftInter[1] {
			dst = append(dst, []int{leftInter[0] - src[isrc][2], leftInter[1] - src[isrc][2], src[isrc][2]})
		}

		if midInter[0] <= midInter[1] {
			dst = append(dst, []int{midInter[0] - src[isrc][2], midInter[1] - src[isrc][2], src[isrc][2] + mappInter[2]})
		}

		if rightInter[0] <= rightInter[1] {
			src[isrc] = []int{rightInter[0] - src[isrc][2], rightInter[1] - src[isrc][2], src[isrc][2]}
			isrc--
		}

		isrc++
	}

	dst = append(dst, src[isrc:]...)

	return dst
}

func solve(sc *util.Scanner) any {
	ints := toInts(strings.Fields(sc.NextLine())[1:])
	seeds := make([][]int, 0, len(ints)/2)
	for i := 0; i < len(ints); i += 2 {
		seeds = append(seeds, []int{ints[i], ints[i] + ints[i+1] - 1, 0})
	}
	sc.NextLine()

	sortIntersFunc := func(inter1, inter2 []int) int {
		return (inter1[0] + inter1[2]) - (inter2[0] + inter2[2])
	}
	slices.SortFunc(seeds, sortIntersFunc)

	sortMappsFunc := func(inter1, inter2 []int) int {
		return inter1[0] - inter2[0]
	}
	mapps := make([][][]int, 0)
	for sc.HasNextLine() {
		sc.NextLine()

		mapp := make([][]int, 0)
		for sc.HasNextLine() {
			ints := toInts(strings.Fields(sc.NextLine()))
			if len(ints) == 0 {
				break
			}

			mapp = append(mapp, []int{ints[1], ints[1] + ints[2] - 1, ints[0] - ints[1]})
			slices.SortFunc(mapp, sortMappsFunc)
		}
		mapps = append(mapps, mapp)
	}

	inters := seeds
	for _, mapp := range mapps {
		inters = applyMapp(inters, mapp)
		slices.SortFunc(inters, sortIntersFunc)
	}

	return inters[0][0] + inters[0][2]
}

func main() {
	util.RunAll("Day 5", solve)
}

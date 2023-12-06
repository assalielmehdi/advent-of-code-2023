package main

import (
	util_fp "assalielmehdi/adventofcode2023/pkg/fp"
	util_io "assalielmehdi/adventofcode2023/pkg/io"
	util_slice "assalielmehdi/adventofcode2023/pkg/slice"
	"fmt"
	"strconv"
	"strings"
)

type Interval struct {
	from  int
	to    int
	delta int
}

type Mapping struct {
	intervals []*Interval
}

func mapValues(values []int, mapping *Mapping) []int {
	mapped := make([]int, 0, len(values))

	for _, value := range values {
		delta := 0

		for _, interval := range mapping.intervals {
			if interval.from <= value && value <= interval.to {
				delta = interval.delta
				break
			}
		}

		mapped = append(mapped, value+delta)
	}

	return mapped
}

func parseSeeds(line []byte) []int {
	scanner := util_io.NewScanner(line)
	seeds := make([]int, 0)

	scanner.Next() // seeds:
	for scanner.HasNext() {
		seeds = append(seeds, scanner.NextInt())
	}

	return seeds
}

func parseMapping(lines [][]byte, i int) (int, *Mapping) {
	intervals := make([]*Interval, 0)

	i++
	for i < len(lines) && len(lines[i]) > 0 {
		tokens := util_fp.Map(strings.Split(string(lines[i]), " "), func(valStr string) int {
			val, _ := strconv.Atoi(valStr)
			return val
		})

		intervals = append(intervals, &Interval{
			from:  tokens[1],
			to:    tokens[1] + tokens[2] - 1,
			delta: tokens[0] - tokens[1],
		})

		i++
	}
	i++

	return i, &Mapping{
		intervals: intervals,
	}
}

func solve1() {
	lines := util_io.ReadLines("input1.txt")

	seeds := parseSeeds(lines[0])
	mappings := make([]*Mapping, 0)

	for i := 2; i < len(lines); {
		j, mapping := parseMapping(lines, i)
		mappings = append(mappings, mapping)
		i = j
	}

	values := util_fp.Reduce(mappings, func(values []int, mapping *Mapping) []int {
		return mapValues(values, mapping)
	}, seeds)

	fmt.Println(util_slice.Min(values))
}

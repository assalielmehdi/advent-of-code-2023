package main

// import (
// 	util_io "assalielmehdi/adventofcode2023/pkg/io"
// )

// func intersection(interval1, interval2 *Interval) *Interval {
// 	return &Interval{
// 		from:  max(interval1.from, interval2.from),
// 		to:    min(interval1.to, interval2.to),
// 		delta: interval1.delta + interval1.delta,
// 	}
// }

// func parseSeedsIntervals(line []byte) []*Interval {
// 	scanner := util_io.NewScanner(line)
// 	seedsIntervals := make([]*Interval, 0)

// 	scanner.Next() // seeds:
// 	for scanner.HasNext() {
// 		from, len := scanner.NextInt(), scanner.NextInt()

// 		seedsIntervals = append(seedsIntervals, &Interval{
// 			from: from,
// 			to:   from + len - 1,
// 		})
// 	}

// 	return seedsIntervals
// }

// func mapIntervals(values []int, mapping *Mapping) []int {
// 	mapped := make([]int, 0, len(values))

// 	for _, value := range values {
// 		delta := 0

// 		for _, interval := range mapping.intervals {
// 			if interval.from <= value && value <= interval.to {
// 				delta = interval.delta
// 				break
// 			}
// 		}

// 		mapped = append(mapped, value+delta)
// 	}

// 	return mapped
// }

// func solve2() {
// 	lines := util_io.ReadLines("input1.txt")

// 	seedsIntervals := parseSeedsIntervals(lines[0])
// 	mappings := make([]*Mapping, 0)

// 	for i := 2; i < len(lines); {
// 		j, mapping := parseMapping(lines, i)
// 		mappings = append(mappings, mapping)
// 		i = j
// 	}

// 	intervals := seedsIntervals
// 	for _, mapping := range mappings {
// 		newIntervals := make([]*Interval, 0)

// 		for _, mappingInterval := range mapping.intervals {
// 			for _, interval := range intervals {
// 				inter := intersection(interval, mappingInterval)

// 				// if inter.from
// 			}
// 		}
// 	}
// }

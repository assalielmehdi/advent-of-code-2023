package main

import (
	util_fp "assalielmehdi/adventofcode2023/pkg/fp"
	util_io "assalielmehdi/adventofcode2023/pkg/io"
	"fmt"
	"math"
)

func solveRace(time, distance float64) int64 {
	t1 := (time - math.Sqrt(time*time-4*distance)) * 0.5
	t2 := (time + math.Sqrt(time*time-4*distance)) * 0.5

	if math.Ceil(t1) == t1 {
		t1 = t1 + 1
	} else {
		t1 = math.Ceil(t1)
	}

	if math.Floor(t2) == t2 {
		t2 = t2 - 1
	} else {
		t2 = math.Floor(t2)
	}

	return int64(t2 - t1 + 1)
}

func main() {
	lines := util_io.ReadLines("input2.txt")

	scanner := util_io.NewScanner(lines[0])
	scanner.Next() // Time:
	times := util_fp.Map(scanner.NextInts(), func(val int) float64 { return float64(val) })

	scanner = util_io.NewScanner(lines[1])
	scanner.Next() // Distance:
	distances := util_fp.Map(scanner.NextInts(), func(val int) float64 { return float64(val) })

	answer := int64(1)

	for i := range times {
		answer *= solveRace(times[i], distances[i])
	}

	fmt.Println(answer)
}

package main

import (
	"fmt"
	"math"

	util_fp "assalielmehdi/adventofcode2023/pkg/fp"
	util_io "assalielmehdi/adventofcode2023/pkg/io"
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

	answer := util_fp.Reduce(
		util_fp.Zip(times, distances),
		func(answer int64, race *util_fp.Pair[float64, float64]) int64 {
			return answer * solveRace(race.First, race.Second)
		},
		int64(1),
	)

	fmt.Println(answer)
}

package main

import (
	"math"
	"strconv"
	"strings"

	"assalielmehdi/adventofcode2023/util"
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

func solve(sc *util.Scanner) any {
	sc.Next() // Time:
	timesStr := strings.Fields(sc.NextLine())
	times := make([]float64, 0, len(timesStr))
	for _, timeStr := range timesStr {
		time, _ := strconv.Atoi(timeStr)
		times = append(times, float64(time))
	}

	sc.Next() // Distance:
	distancesStr := strings.Fields(sc.NextLine())
	distances := make([]float64, 0, len(distancesStr))
	for _, distanceStr := range distancesStr {
		distance, _ := strconv.Atoi(distanceStr)
		distances = append(distances, float64(distance))
	}

	answer := int64(1)

	for i := range times {
		answer *= solveRace(times[i], distances[i])
	}

	return answer
}

func main() {
	util.RunAll("Day 6", solve)
}

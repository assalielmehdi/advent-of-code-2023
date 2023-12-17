package main

import (
	"slices"
	"strconv"
	"strings"

	"assalielmehdi/adventofcode2023/util"
)

type lens struct {
	name string
	len  int
}

func solve1(sc *util.Scanner) any {
	tokens := strings.Split(sc.NextLine(), ",")
	sum := 0
	for _, token := range tokens {
		hash := 0
		for _, ch := range token {
			hash = ((hash + int(ch)) * 17) % 256
		}
		sum += hash
	}
	return sum
}

func hash(s string) int {
	h := 0
	for _, ch := range s {
		h = ((h + int(ch)) * 17) % 256
	}
	return h
}

func remStep(boxes [][]lens, name string) {
	h := hash(name)
	idx := slices.IndexFunc(boxes[h], func(l lens) bool {
		return l.name == name
	})
	if idx == -1 {
		return
	}
	boxes[h] = append(boxes[h][:idx], boxes[h][idx+1:]...)
}

func addStep(boxes [][]lens, l lens) {
	h := hash(l.name)
	idx := slices.IndexFunc(boxes[h], func(ll lens) bool {
		return ll.name == l.name
	})
	if idx == -1 {
		boxes[h] = append(boxes[h], l)
		return
	}
	boxes[h][idx].len = l.len
}

func solve2(sc *util.Scanner) any {
	tokens := strings.Split(sc.NextLine(), ",")
	boxes := make([][]lens, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = make([]lens, 0)
	}
	for _, token := range tokens {
		if strings.HasSuffix(token, "-") {
			box, _ := strings.CutSuffix(token, "-")
			remStep(boxes, box)
		} else {
			box := strings.Split(token, "=")
			len, _ := strconv.Atoi(box[1])
			addStep(boxes, lens{box[0], len})
		}
	}
	answer := 0
	for i, box := range boxes {
		for j, l := range box {
			answer += (i + 1) * (j + 1) * l.len
		}
	}
	return answer
}

func main() {
	util.RunAll("Day 15 - Part 1", solve1)
	util.RunAll("Day 15 - Part 2", solve2)
}

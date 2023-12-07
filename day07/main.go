package main

import util_map "assalielmehdi/adventofcode2023/pkg/map"

type Hand struct {
	Cards []byte
	Type  int
	Bid   int
}

func typeOfHand(cards []byte) int {
	occ := make(map[byte]int)

	for _, card := range cards {
		if card == 'J' {
			continue
		}

		occ[card]++
	}

	if util_map.CountVal(occ, 5) == 1 || len(occ) == 1 { // Five of a kin
		return 6
	}

	if util_map.CountVal(occ, 4) == 1 && util_map.CountVal(occ, 1) == 1 { // Four of a kind
		return 5
	}

	if util_map.CountVal(occ, 3) == 1 && util_map.CountVal(occ, 2) == 1 { // Full house
		return 4
	}

	if util_map.CountVal(occ, 3) == 1 && util_map.CountVal(occ, 1) == 2 { // Three of a kind
		return 3
	}

	if util_map.CountVal(occ, 2) == 2 && util_map.CountVal(occ, 1) == 1 { // Two pair
		return 2
	}

	if util_map.CountVal(occ, 2) == 1 && util_map.CountVal(occ, 1) == 3 { // One pair
		return 1
	}

	return 0
}

func main() {
	solve1()
	solve2()
}

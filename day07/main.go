package main

import (
	"assalielmehdi/adventofcode2023/util"
	"slices"
)

type Hand struct {
	Cards []byte
	Type  int
	Bid   int
}

func countVal(vals map[byte]int, search int) int {
	count := 0

	for _, val := range vals {
		if val == search {
			count++
		}
	}

	return count
}

func typeOfHand(cards []byte) int {
	occ := make(map[byte]int)

	for _, card := range cards {
		if card == 'J' {
			continue
		}

		occ[card]++
	}

	if countVal(occ, 5) == 1 || len(occ) == 1 { // Five of a kin
		return 6
	}

	if countVal(occ, 4) == 1 && countVal(occ, 1) == 1 { // Four of a kind
		return 5
	}

	if countVal(occ, 3) == 1 && countVal(occ, 2) == 1 { // Full house
		return 4
	}

	if countVal(occ, 3) == 1 && countVal(occ, 1) == 2 { // Three of a kind
		return 3
	}

	if countVal(occ, 2) == 2 && countVal(occ, 1) == 1 { // Two pair
		return 2
	}

	if countVal(occ, 2) == 1 && countVal(occ, 1) == 3 { // One pair
		return 1
	}

	return 0
}

func NewHand1(cards []byte, bid int) *Hand {
	return &Hand{
		Cards: cards,
		Type:  typeOfHand(cards),
		Bid:   bid,
	}
}

func compareHands1(h1, h2 *Hand) int {
	if h1.Type != h2.Type {
		return h1.Type - h2.Type
	}

	cardsOrder := []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

	for i := range h1.Cards {
		order1 := slices.Index(cardsOrder, h1.Cards[i])
		order2 := slices.Index(cardsOrder, h2.Cards[i])

		if order1 != order2 {
			return order1 - order2
		}
	}

	return 0
}

func solve1(sc *util.Scanner) any {
	hands := make([]*Hand, 0)

	for sc.HasNext() {
		hands = append(hands, NewHand1(sc.NextBytes(), sc.NextInt()))
	}

	slices.SortFunc(hands, compareHands1)

	score := 0

	for i, hand := range hands {
		score += hand.Bid * (i + 1)
	}

	return score
}

func maxPossibleTypeOfHand(cards []byte, i int) int {
	if i == len(cards) {
		return typeOfHand(cards)
	}

	if cards[i] != 'J' {
		return maxPossibleTypeOfHand(cards, i+1)
	}

	sortedCards := []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
	maxPossibleType := 0

	for _, card := range sortedCards {
		if card == 'J' {
			continue
		}

		cards[i] = card
		maxPossibleType = max(maxPossibleType, maxPossibleTypeOfHand(cards, i+1))
		cards[i] = 'J'
	}

	return maxPossibleType
}

func NewHand2(cards []byte, bid int) *Hand {
	return &Hand{
		Cards: cards,
		Type:  maxPossibleTypeOfHand(cards, 0),
		Bid:   bid,
	}
}

func compareHands2(h1, h2 *Hand) int {
	if h1.Type != h2.Type {
		return h1.Type - h2.Type
	}

	sortedCards := []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

	for i := range h1.Cards {
		order1 := slices.Index(sortedCards, h1.Cards[i])
		order2 := slices.Index(sortedCards, h2.Cards[i])

		if order1 != order2 {
			return order1 - order2
		}
	}

	return 0
}

func solve2(sc *util.Scanner) any {
	hands := make([]*Hand, 0)

	for sc.HasNext() {
		hands = append(hands, NewHand2(sc.NextBytes(), sc.NextInt()))
	}

	slices.SortFunc(hands, compareHands2)

	score := 0

	for i, hand := range hands {
		score += hand.Bid * (i + 1)
	}

	return score
}

func main() {
	util.RunAll("Day 7 - Part 1", solve1)
	util.RunAll("Day 7 - Part 2", solve2)
}

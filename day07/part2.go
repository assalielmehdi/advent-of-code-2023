package main

import (
	"fmt"
	"slices"

	util_io "assalielmehdi/adventofcode2023/pkg/io"
)

var sortedCards = []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

func maxPossibleTypeOfHand(cards []byte, i int) int {
	if i == len(cards) {
		return typeOfHand(cards)
	}

	if cards[i] != 'J' {
		return maxPossibleTypeOfHand(cards, i+1)
	}

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

	for i := range h1.Cards {
		order1 := slices.Index(sortedCards, h1.Cards[i])
		order2 := slices.Index(sortedCards, h2.Cards[i])

		if order1 != order2 {
			return order1 - order2
		}
	}

	return 0
}

func solve2() {
	lines := util_io.ReadLines("input1.txt")
	scanner := util_io.NewEmptyScanner()
	hands := make([]*Hand, 0, len(lines))

	for _, line := range lines {
		scanner.SetContent(line)
		hands = append(hands, NewHand2(scanner.Next(), scanner.NextInt()))
	}

	slices.SortFunc(hands, compareHands2)

	score := 0

	for i, hand := range hands {
		score += hand.Bid * (i + 1)
	}

	fmt.Println(score)
}

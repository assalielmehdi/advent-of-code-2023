package main

import (
	"fmt"
	"slices"

	util_io "assalielmehdi/adventofcode2023/pkg/io"
)

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

func solve1() {
	lines := util_io.ReadLines("input1.txt")
	scanner := util_io.NewEmptyScanner()
	hands := make([]*Hand, 0, len(lines))

	for _, line := range lines {
		scanner.SetContent(line)
		hands = append(hands, NewHand1(scanner.Next(), scanner.NextInt()))
	}

	slices.SortFunc(hands, compareHands1)

	score := 0

	for i, hand := range hands {
		score += hand.Bid * (i + 1)
	}

	fmt.Println(score)
}

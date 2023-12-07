// Package main for Advent of Code 2023, day 7, part 1
// https://adventofcode.com/2023/day/7
package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Card rune

const (
	Err   Card = '!'
	Empty Card = '~'
)

type HandType int

const (
	HighCard HandType = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var mapOfHandType = map[HandType]string{
	1: "HighCard",
	2: "OnePair",
	3: "TwoPair",
	4: "ThreeOfAKind",
	5: "FullHouse",
	6: "FourOfAKind",
	7: "FiveOfAKind",
}

var strengthRankings = []Card{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

func (c Card) StrongerThan(d Card) bool {
	if d == Empty {
		return true
	}

	cIndex := slices.Index(strengthRankings, c)
	dIndex := slices.Index(strengthRankings, d)

	return cIndex < dIndex
}

type Hand struct {
	cards    [5]Card
	handType HandType
}

/*
StrongerThan implements this section:

	Five of a kind, where all five cards have the same label: AAAAA
	Four of a kind, where four cards have the same label and one card has a different label: AA8AA
	Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
	Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
	Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
	One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
	High card, where all cards' labels are distinct: 23456
*/
func (h Hand) StrongerThan(i Hand) Hand {
	if h.handType > i.handType {
		return h
	}

	if h.handType < i.handType {
		return i
	}

	if h.secondOrderingRule(i) {
		return h
	} else {
		return i
	}
}

// xOfAKind asserts whether this hand is x-of-a-kind or not.
// If you're looking for a different -of-a-kind to a particular card, supply that, otherwise give the Empty Card.
func (h Hand) xOfAKind(x int, not Card) (bool, Card) {
	m := map[Card]int{}

	for i := 0; i < len(h.cards); i++ {
		m[h.cards[i]]++
	}

	for c, b := range m {
		if b == x && c != not {
			return true, c
		}
	}

	return false, Err
}

// fullHouse looks for differing three- and two-of-a-kind.
func (h Hand) fullHouse() bool {
	threeResult, threeCard := h.xOfAKind(3, Empty)
	twoResult, _ := h.xOfAKind(2, threeCard)

	return threeResult && twoResult
}

// twoPair returns the card of each pair, and the remaining card last.
func (h Hand) twoPair() (bool, Card, Card, Card) {
	firstResult, firstCard := h.xOfAKind(2, Empty)
	secondResult, secondCard := h.xOfAKind(2, firstCard)

	remaining := Empty

	for i := 0; i < len(h.cards); i++ {
		if h.cards[i] != firstCard && h.cards[i] != secondCard {
			remaining = h.cards[i]
			break
		}
	}

	return firstResult && secondResult && firstCard != secondCard && firstCard != Err && secondCard != Err,
		firstCard, secondCard, remaining
}

// onePair returns the card of the pair, and the remaining three cards in hand order.
func (h Hand) onePair() (bool, Card, Card, Card, Card) {
	firstResult, firstCard := h.xOfAKind(2, Empty)
	if !firstResult {
		return false, Err, Err, Err, Err
	}

	m := map[int][]Card{}

	for i := 0; i < len(h.cards); i++ {
		card := h.cards[i]
		key := h.count(card)

		if _, exists := m[key]; !exists {
			m[key] = make([]Card, 0)
		}

		m[key] = append(m[key], card)
	}

	// If there aren't 3 instances of cards each with a count of 1 in this hand...
	if len(m[1]) != 3 {
		return false, Err, Err, Err, Err
	}

	return true, firstCard, m[1][0], m[1][1], m[1][2]
}

// highCard returns the high card of the hand, or Err if there is more than one instance of any card in the hand.
func (h Hand) highCard() Card {
	highest := Empty

	for i := 0; i < len(h.cards); i++ {
		for j := i + 1; j < len(h.cards); j++ {
			if h.cards[i] == h.cards[j] {
				return Err
			}
		}

		if h.cards[i].StrongerThan(highest) {
			highest = h.cards[i]
		}
	}

	return highest
}

/*
secondOrderingRule implements this section:

	If two hands have the same type, a second ordering rule takes effect. Start by comparing the first card in each hand. If these cards are different, the hand with the stronger first card is considered stronger. If the first card in each hand have the same label, however, then move on to considering the second card in each hand. If they differ, the hand with the higher second card wins; otherwise, continue with the third card in each hand, then the fourth, then the fifth.
*/
func (h Hand) secondOrderingRule(i Hand) bool {
	for j := 0; j < len(h.cards); j++ {
		if h.cards[j] == i.cards[j] {
			continue
		}

		return h.cards[j].StrongerThan(i.cards[j])
	}

	return false
}

// count how many of this Card in this Hand.
func (h Hand) count(c Card) int {
	return strings.Count(h.String(), string(c))
}

func (h Hand) String() string {
	return fmt.Sprintf("%s%s%s%s%s (%s)",
		string(h.cards[0]), string(h.cards[1]), string(h.cards[2]), string(h.cards[3]), string(h.cards[4]),
		mapOfHandType[h.handType])
}

var ErrZeroLengthLine = errors.New("line has length of zero")

func parseLine(input string) (Hand, int, error) {
	if len(input) == 0 {
		return Hand{}, 0, ErrZeroLengthLine
	}

	xInput := strings.Split(strings.TrimSpace(input), " ")
	if len(xInput) != 2 {
		return Hand{}, 0, fmt.Errorf("not two fields on line '%s'", input)
	}

	hand, err := parseHand(xInput[0])
	if err != nil {
		return Hand{}, 0, err
	}

	bid, err := strconv.Atoi(xInput[1])
	if err != nil {
		return Hand{}, 0, err
	}

	return hand, bid, nil
}

func parseHand(input string) (Hand, error) {
	var h Hand

	if len(input) != 5 {
		return Hand{}, fmt.Errorf("input '%s' should have 5 cards exactly", input)
	}

	for i := 0; i < 5; i++ {
		h.cards[i] = []Card(input)[i]
	}

	if h5, _ := h.xOfAKind(5, Empty); h5 {
		h.handType = FiveOfAKind
	} else if h4, _ := h.xOfAKind(4, Empty); h4 {
		h.handType = FourOfAKind
	} else if h.fullHouse() {
		h.handType = FullHouse
	} else if h3, _ := h.xOfAKind(3, Empty); h3 {
		h.handType = ThreeOfAKind
	} else if h2, _, _, _ := h.twoPair(); h2 {
		h.handType = TwoPair
	} else if h1, _, _, _, _ := h.onePair(); h1 {
		h.handType = OnePair
	} else {
		h.handType = HighCard
	}

	return h, nil
}

func Part1(inputLines []string) (int, error) {
	bids := map[Hand]int{}

	for ilIndex := range inputLines {
		hand, bid, err := parseLine(inputLines[ilIndex])
		if err != nil {
			if errors.Is(err, ErrZeroLengthLine) {
				continue
			}

			return 0, err
		}

		if _, exists := bids[hand]; exists {
			// watch out
			return 0, fmt.Errorf("hand '%s' already exists in deck", hand)
		}

		bids[hand] = bid
	}

	rankingOfHands := []Hand{}

	for hand := range bids {
		rankingOfHands = append(rankingOfHands, hand)
	}

	slices.SortStableFunc(rankingOfHands, func(a, b Hand) int {
		result := a.StrongerThan(b)

		if result == a {
			return 1
		} else {
			return -1
		}
	})

	totalWinnings := 0

	for rank, hand := range rankingOfHands {
		rank += 1

		bid := bids[hand]

		totalWinnings += rank * bid
	}

	return totalWinnings, nil
}

// Package main for Advent of Code 2023, day 4, part 2
// https://adventofcode.com/2023/day/4
package main

import (
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"
)

func Part2(inputLines []string) (int, error) {
	cardDeck := make(Cards, 0)

	for _, inputLine := range inputLines {
		if len(inputLine) == 0 {
			continue
		}

		slog.Debug("input card",
			slog.String("line", inputLine))

		lineCard := Card{}

		xCardNum := strings.Split(strings.TrimSpace(inputLine), ":")
		if len(xCardNum) != 2 {
			return 0, nil
		}

		cardNumOnly := strings.TrimPrefix(xCardNum[0], "Card ")

		cardNum, err := strconv.Atoi(strings.TrimSpace(cardNumOnly))
		if err != nil {
			return 0, fmt.Errorf("first strconv.Atoi: %w", err)
		}

		lineCard.number = cardNum

		xNumbers := strings.Split(strings.TrimSpace(xCardNum[1]), "|")
		if len(xNumbers) != 2 {
			return 0, nil
		}

		winningNumbers := strings.Split(strings.TrimSpace(xNumbers[0]), " ")
		xwn := make([]int, 0)

		for _, wn := range winningNumbers {
			if wn == "" {
				continue
			}

			iwn, err := strconv.Atoi(wn)
			if err != nil {
				return 0, err
			}

			xwn = append(xwn, iwn)
		}

		lineCard.winningNumbers = xwn

		numbersYouHave := strings.Split(strings.TrimSpace(xNumbers[1]), " ")
		xnyh := make([]int, 0)

		for _, nyh := range numbersYouHave {
			if nyh == "" {
				continue
			}

			inyh, err := strconv.Atoi(nyh)
			if err != nil {
				return 0, err
			}

			xnyh = append(xnyh, inyh)
		}

		lineCard.numbersYouHave = xnyh

		cardDeck = append(cardDeck, lineCard)
	}

	slog.Info("parsed deck",
		slog.Any("card_deck", cardDeck))

	cardIndex := 0

	// cardNumbers maps card number to number of occurences, both from the original, as well as won copies.
	cardNumbers := map[int]int{}

	// Set up one each for the original cards.
	for cardIndex < len(cardDeck) {
		cardNumbers[cardDeck[cardIndex].number] = 1
		cardIndex++
	}

	cardIndex = 0

	// For each card that wins, add copies of subsequent cards to the above map.
	for cardIndex < len(cardDeck) {
		winsFromCard := cardDeck[cardIndex].CountWins()

		copiesOfWinningCard := cardNumbers[cardDeck[cardIndex].number]

		slog.Info("wins from card",
			slog.Int("card_number", cardDeck[cardIndex].number),
			slog.Int("wins_from_card", winsFromCard),
			slog.Int("have_copies_of_winning_card", copiesOfWinningCard))

		for i := 1; i <= winsFromCard; i++ {
			wonACopy := cardDeck[cardIndex].number + i

			slog.Info("won a copy",
				slog.Int("win_came_from", cardDeck[cardIndex].number),
				slog.Int("win_copy_of", wonACopy))

			cardNumbers[wonACopy] += copiesOfWinningCard
		}

		slog.Info("running tally of copies won",
			slog.Any("deck", cardNumbers))

		cardIndex++
	}

	slog.Info("card numbers",
		slog.Any("map", cardNumbers))

	// Tally up the counts in the map.
	countOccurrences := 0
	for _, count := range cardNumbers {
		countOccurrences += count
	}

	slog.Info("occurrences",
		slog.Any("count_occurrences", countOccurrences))

	return countOccurrences, nil
}

type Card struct {
	number                         int
	winningNumbers, numbersYouHave []int
}

func (c Card) CountWins() int {
	wins := 0

	for _, nyh := range c.numbersYouHave {
		if slices.Contains(c.winningNumbers, nyh) {
			wins++
		}
	}

	return wins
}

type Cards []Card

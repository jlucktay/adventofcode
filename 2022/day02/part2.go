package day02

import (
	"bufio"
	"fmt"
	"strings"
	"unicode/utf8"
)

func StrategisedScore(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		shapes := strings.Split(line, " ")

		if len(shapes) != 2 {
			return 0, fmt.Errorf("could not split line '%s' on a space", line)
		}

		rOpp, _ := utf8.DecodeRuneInString(shapes[0])
		rPlay, _ := utf8.DecodeRuneInString(shapes[1])

		if rOpp == utf8.RuneError || rPlay == utf8.RuneError {
			return 0, fmt.Errorf("could not decode rune from '%v'", shapes)
		}

		total += strategiseRPS(rOpp, rPlay)
	}

	return total, nil
}

// Player
// X means you need to lose
// Y means you need to end the round in a draw
// Z means you need to win
//
// (1 for Rock, 2 for Paper, and 3 for Scissors)
//
// Opponent
// A for Rock, B for Paper, and C for Scissors
// (0 if you lost, 3 if the round was a draw, and 6 if you won)

func strategiseRPS(opponent, player rune) int {
	score := 0

	// Player choice points
	switch player {
	case 'X':
		switch opponent {
		case 'A':
			score += 3
		case 'B':
			score += 1
		case 'C':
			score += 2
		}
	case 'Y':
		switch opponent {
		case 'A':
			score += 1
		case 'B':
			score += 2
		case 'C':
			score += 3
		}
	case 'Z':
		switch opponent {
		case 'A':
			score += 2
		case 'B':
			score += 3
		case 'C':
			score += 1
		}
	}

	// Win/lose/draw points
	switch player {
	case 'X':
		// noop
	case 'Y':
		score += 3
	case 'Z':
		score += 6
	}

	return score
}

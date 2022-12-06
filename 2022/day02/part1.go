package day02

import (
	"bufio"
	"fmt"
	"strings"
	"unicode/utf8"
)

func TotalScore(input string) (int, error) {
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

		total += playRPS(rOpp, rPlay)
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanning input: %v", err)
	}

	return total, nil
}

// Player
// X for Rock, Y for Paper, and Z for Scissors
// (1 for Rock, 2 for Paper, and 3 for Scissors)
//
// Opponent
// A for Rock, B for Paper, and C for Scissors
// (0 if you lost, 3 if the round was a draw, and 6 if you won)

func playRPS(opponent, player rune) int {
	score := 0

	switch player {
	case 'X':
		score += 1
	case 'Y':
		score += 2
	case 'Z':
		score += 3
	}

	switch opponent {
	case 'A':
		switch player {
		case 'X':
			score += 3
		case 'Y':
			score += 6
		case 'Z':
			// noop
		}
	case 'B':
		switch player {
		case 'X':
			// noop
		case 'Y':
			score += 3
		case 'Z':
			score += 6
		}
	case 'C':
		switch player {
		case 'X':
			score += 6
		case 'Y':
			// noop
		case 'Z':
			score += 3
		}
	}

	return score
}

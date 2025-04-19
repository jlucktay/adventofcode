package day02

import (
	"bufio"
	"fmt"
	"strings"
	"unicode/utf8"
)

type rpsStrategy func(rune, rune) int

func calculateScore(strat rpsStrategy, input string) (int, error) {
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

		total += strat(rOpp, rPlay)
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanning input: %w", err)
	}

	return total, nil
}

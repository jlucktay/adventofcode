package day05

import (
	"bufio"
	"fmt"
	"strings"
)

func TopCrate9001(input string) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	afterSeperator := false

	crateDeques := make([]crateDeque, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			afterSeperator = true
			continue
		}

		if !afterSeperator {
			if err := parseLineOfCrates(&crateDeques, line); err != nil {
				return "", fmt.Errorf("parsing crates '%s': %w", line, err)
			}
		} else {
			// Below the seperator, parse 'move X from Y to Z'
			qty, from, to, err := parseMoveOrders(&crateDeques, line)
			if err != nil {
				return "", fmt.Errorf("parsing move orders '%s': %w", line, err)
			}

			if err := crateMover9001(&crateDeques, qty, from, to); err != nil {
				return "", fmt.Errorf("moving crates with 9000: %w", err)
			}
		}
	}

	finalOrder := ""
	for i := 0; i < len(crateDeques); i++ {
		firstCrate, ok := crateDeques[i].getFirst()
		if !ok {
			return "", fmt.Errorf("getting first from stack #%d", i)
		}

		finalOrder += string(firstCrate)
	}

	return finalOrder, nil
}

func crateMover9001(stacks *[]crateDeque, qty, from, to int) error {
	moveTheseCrates := strings.Builder{}

	for i := 0; i < qty; i++ {
		crate, ok := (*stacks)[from-1].popFirst()
		if !ok {
			return fmt.Errorf("getting first from stack #%d", from)
		}

		moveTheseCrates.WriteString(string(crate))
	}

	mtc := moveTheseCrates.String()

	for j := len(mtc) - 1; j >= 0; j-- {
		(*stacks)[to-1].prepend(rune(mtc[j]))
	}

	return nil
}

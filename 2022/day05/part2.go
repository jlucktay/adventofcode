package day05

import (
	"bufio"
	"container/list"
	"fmt"
	"strings"
)

func TopCrate9001(input string) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	afterSeperator := false

	crateStacks := make([]*list.List, 0)

	var err error

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			afterSeperator = true

			continue
		}

		if !afterSeperator {
			if crateStacks, err = parseLineOfCrates(crateStacks, line); err != nil {
				return "", fmt.Errorf("parsing crates '%s': %w", line, err)
			}
		} else {
			// Below the separator, parse 'move X from Y to Z'
			qty, from, to, err := parseMoveOrders(line)
			if err != nil {
				return "", fmt.Errorf("parsing move orders '%s': %w", line, err)
			}

			if crateStacks, err = crateMover9001(crateStacks, qty, from, to); err != nil {
				return "", fmt.Errorf("moving crates with 9000: %w", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanning input: %w", err)
	}

	finalOrder := ""

	for i := range len(crateStacks) {
		firstCrate := crateStacks[i].Front()
		if firstCrate == nil {
			return "", fmt.Errorf("getting first from stack #%d", i)
		}

		fcv, ok := firstCrate.Value.(rune)
		if !ok {
			return "", fmt.Errorf("type asserting on first from stack #%d", i)
		}

		finalOrder += string(fcv)
	}

	return finalOrder, nil
}

func crateMover9001(stacks []*list.List, qty, from, to int) ([]*list.List, error) {
	moveTheseCrates := list.New()

	for i := range qty {
		crate := stacks[from-1].Front()
		if crate == nil {
			return nil, fmt.Errorf("getting first from stack #%d", from)
		}

		cvr, ok := crate.Value.(rune)
		if !ok {
			return nil, fmt.Errorf("type asserting on first from stack #%d", i)
		}

		stacks[from-1].Remove(crate)
		moveTheseCrates.PushBack(cvr)
	}

	stacks[to-1].PushFrontList(moveTheseCrates)

	return stacks, nil
}

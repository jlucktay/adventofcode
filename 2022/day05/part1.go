package day05

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func TopCrate9000(input string) (string, error) {
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
			// Below the seperator, parse 'move X from Y to Z'
			qty, from, to, err := parseMoveOrders(line)
			if err != nil {
				return "", fmt.Errorf("parsing move orders '%s': %w", line, err)
			}

			if crateStacks, err = crateMover9000(crateStacks, qty, from, to); err != nil {
				return "", fmt.Errorf("moving crates with 9000: %w", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanning input: %v", err)
	}

	finalOrder := ""

	for i := 0; i < len(crateStacks); i++ {
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

func parseLineOfCrates(incoming []*list.List, line string) ([]*list.List, error) {
	lineSegment := line[0:3]
	crateStackIndex := 0
	lineOffset := 0

	for {
		trimmed := strings.Trim(lineSegment, " []")

		if len(trimmed) > 0 {
			r, _ := utf8.DecodeRuneInString(trimmed)
			if r == utf8.RuneError {
				return nil, fmt.Errorf("decoding rune from '%v'", trimmed)
			}

			for len(incoming) < crateStackIndex+1 {
				incoming = append(incoming, list.New())
			}

			if unicode.IsLetter(r) {
				incoming[crateStackIndex].PushBack(r)
			} else if unicode.IsNumber(r) {
				break
			}
		}

		crateStackIndex++

		if len(line) <= len(lineSegment)+lineOffset+1 {
			break
		}

		lineOffset += 3
		lineSegment = line[lineOffset : lineOffset+3]
		lineOffset++
	}

	return incoming, nil
}

func parseMoveOrders(orders string) (int, int, int, error) {
	pattern := `^move ([0-9]+) from ([0-9]+) to ([0-9]+)$`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("compiling regular expression '%s': %w", pattern, err)
	}

	found := regex.FindStringSubmatch(orders)
	if found == nil {
		return 0, 0, 0, fmt.Errorf("finding submatches in '%s'", orders)
	}

	if len(found) != 4 {
		return 0, 0, 0, errors.New("regex SNAFU")
	}

	moveQuantity, err := strconv.Atoi(found[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("parsing int from '%s': %w", found[1], err)
	}

	fromStack, err := strconv.Atoi(found[2])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("parsing int from '%s': %w", found[2], err)
	}

	toStack, err := strconv.Atoi(found[3])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("parsing int from '%s': %w", found[3], err)
	}

	return moveQuantity, fromStack, toStack, nil
}

func crateMover9000(stacks []*list.List, qty, from, to int) ([]*list.List, error) {
	for i := 0; i < qty; i++ {
		crate := stacks[from-1].Front()
		if crate == nil {
			return nil, fmt.Errorf("getting first from stack #%d", from)
		}

		stacks[from-1].Remove(crate)
		stacks[to-1].PushFront(crate.Value)
	}

	return stacks, nil
}

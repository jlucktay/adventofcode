package day05

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func TopCrateOnEachStack(input string) (string, error) {
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
			if err := parseMoveOrders(&crateDeques, line); err != nil {
				return "", fmt.Errorf("parsing move orders '%s': %w", line, err)
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

func parseLineOfCrates(incoming *[]crateDeque, line string) error {
	lineSegment := line[0:3]
	crateStackIndex := 0
	lineOffset := 0

	for {
		trimmed := strings.Trim(lineSegment, " []")

		if len(trimmed) > 0 {
			r, _ := utf8.DecodeRuneInString(trimmed)
			if r == utf8.RuneError {
				return fmt.Errorf("decoding rune from '%v'", trimmed)
			}

			for len(*incoming) < crateStackIndex+1 {
				*incoming = append(*incoming, make(crateDeque, 0))
			}

			if unicode.IsLetter(r) {
				(*incoming)[crateStackIndex].append(r)
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

	return nil
}

func parseMoveOrders(crates *[]crateDeque, orders string) error {
	pattern := `^move ([0-9]+) from ([0-9]+) to ([0-9]+)$`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("compiling regular expression '%s': %w", pattern, err)
	}

	found := regex.FindStringSubmatch(orders)
	if found == nil {
		return fmt.Errorf("finding submatches in '%s'", orders)
	}

	if len(found) != 4 {
		return errors.New("regex SNAFU")
	}

	moveQuantity, err := strconv.ParseInt(found[1], 10, 32)
	if err != nil {
		return fmt.Errorf("parsing int from '%s': %w", found[1], err)
	}

	fromStack, err := strconv.ParseInt(found[2], 10, 32)
	if err != nil {
		return fmt.Errorf("parsing int from '%s': %w", found[2], err)
	}

	toStack, err := strconv.ParseInt(found[3], 10, 32)
	if err != nil {
		return fmt.Errorf("parsing int from '%s': %w", found[3], err)
	}

	for i := int64(0); i < moveQuantity; i++ {
		crate, ok := (*crates)[fromStack-1].popFirst()
		if !ok {
			return fmt.Errorf("getting first from stack #%d", fromStack)
		}

		(*crates)[toStack-1].prepend(crate)
	}

	return nil
}

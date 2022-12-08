package day08

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type TallTreeGrid struct {
	// First index is vertical, second index is horizontal.
	trees [][]int
}

func NewTallTreeGrid(input string) (TallTreeGrid, error) {
	ttg := TallTreeGrid{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	lineCounter := 0

	for scanner.Scan() {
		line := scanner.Text()

		ttg.trees = append(ttg.trees, make([]int, 0))
		ttg.trees[lineCounter] = make([]int, len(line))

		for i := 0; i < len(line); i++ {
			r, _ := utf8.DecodeRuneInString(line[i:])
			if r == utf8.RuneError {
				return TallTreeGrid{}, fmt.Errorf("decoding rune at index '%d' from '%s'", lineCounter, line)
			}

			treeHeight, err := strconv.ParseInt(string(r), 10, 32)
			if err != nil {
				return TallTreeGrid{}, fmt.Errorf("parsing height from '%s': %w", string(r), err)
			}

			ttg.trees[lineCounter][i] = int(treeHeight)
		}

		lineCounter++
	}

	if err := scanner.Err(); err != nil {
		return TallTreeGrid{}, fmt.Errorf("scanning input: %v", err)
	}

	return ttg, nil
}

package day03

import (
	"bufio"
	"fmt"
	"strings"
	"unicode/utf8"
)

func SplitThreeLines(data []byte, atEOF bool) (int, []byte, error) {
	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	advance, newlines := 0, 0
	r, size := utf8.DecodeRune(data)

	for ; size > 0; r, size = utf8.DecodeRune(data[advance:]) {
		advance += size

		if r == '\n' {
			newlines++
		}

		if newlines >= 3 {
			return advance, data[:advance], nil
		}
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func RucksackGroupPriority(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(SplitThreeLines)

	total := 0

	for scanner.Scan() {
		scanned := scanner.Text()
		xScanned := strings.Split(scanned, "\n")

		if len(xScanned) != 4 {
			return 0, fmt.Errorf("scanning exactly three lines (all with trailing newlines) from '%s'", scanned)
		}

		one := xScanned[0]
		two := xScanned[1]
		three := xScanned[2]

		rOne, err := getRunes(one)
		if err != nil {
			return 0, err
		}

		rTwo, err := getRunes(two)
		if err != nil {
			return 0, err
		}

		rThree, err := getRunes(three)
		if err != nil {
			return 0, err
		}

		badge := findBadge(rOne, rTwo, rThree)
		total += toNum(badge)
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanning input: %w", err)
	}

	return total, nil
}

func findBadge(one, two, three map[rune]struct{}) rune {
	for k := range one {
		_, twoOK := two[k]
		_, threeOK := three[k]

		if twoOK && threeOK {
			return k
		}
	}

	return ' '
}

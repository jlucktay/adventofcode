package day03

import (
	"bufio"
	"fmt"
	"strings"
)

func RucksackPrioritySum(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line)%2 != 0 {
			return 0, fmt.Errorf("uneven number of characters on line: '%s'", line)
		}

		left := line[0 : len(line)/2]
		right := line[len(line)/2:]

		rLeft, err := getRunes(left)
		if err != nil {
			return 0, err
		}

		rRight, err := getRunes(right)
		if err != nil {
			return 0, err
		}

		cr := commonRune(rLeft, rRight)
		total += toNum(cr)
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scanning input: %w", err)
	}

	return total, nil
}

func commonRune(left, right map[rune]struct{}) rune {
	for k := range left {
		if _, ok := right[k]; ok {
			return k
		}
	}

	return ' '
}

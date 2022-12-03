package day03

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
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

	return total, nil
}

func getRunes(rucksackHalf string) (map[rune]struct{}, error) {
	if strings.Contains(rucksackHalf, "\n") {
		return nil, fmt.Errorf("newline in rucksack half '%s'", rucksackHalf)
	}

	items := map[rune]struct{}{}

	for _, runeValue := range rucksackHalf {
		items[runeValue] = struct{}{}
	}

	return items, nil
}

func commonRune(left, right map[rune]struct{}) rune {
	for k := range left {
		if _, ok := right[k]; ok {
			return k
		}
	}

	return ' '
}

func toNum(r rune) int {
	if unicode.IsUpper(r) {
		return int(r%65) + 27
	}

	if unicode.IsLower(r) {
		return int(r%97) + 1
	}

	return 0
}

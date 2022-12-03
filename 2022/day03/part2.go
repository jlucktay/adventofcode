package day03

import (
	"bufio"
	"strings"
)

func RucksackGroupPriority(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	total := 0

	for {
		if !scanner.Scan() {
			break
		}
		one := scanner.Text()

		if !scanner.Scan() {
			break
		}
		two := scanner.Text()

		if !scanner.Scan() {
			break
		}
		three := scanner.Text()

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

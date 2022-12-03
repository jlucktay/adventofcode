package day03

import (
	"fmt"
	"strings"
	"unicode"
)

func getRunes(rucksack string) (map[rune]struct{}, error) {
	if strings.Contains(rucksack, "\n") {
		return nil, fmt.Errorf("newline in rucksack '%s'", rucksack)
	}

	items := map[rune]struct{}{}

	for _, runeValue := range rucksack {
		items[runeValue] = struct{}{}
	}

	return items, nil
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

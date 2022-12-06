package day06

import (
	"fmt"
)

func FindMarkerInDatastream(input string) (int, error) {
	foundChars := [4]rune{}

	for index, r := range input {
		foundChars[index%4] = r
		mapOfArray := make(map[rune]int)

		for _, r := range foundChars {
			// Don't map nulls
			if r > 0 {
				mapOfArray[r]++
			}
		}

		if len(mapOfArray) == 4 {
			return index + 1, nil
		}
	}

	return 0, fmt.Errorf("finding marker in '%s'", input)
}

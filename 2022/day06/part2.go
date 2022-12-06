package day06

import "fmt"

const messageSize = 14

func FindStartOfMessageMarker(input string) (int, error) {
	foundChars := [messageSize]rune{}

	for index, r := range input {
		foundChars[index%messageSize] = r
		mapOfArray := make(map[rune]int)

		for _, r := range foundChars {
			// Don't map nulls
			if r > 0 {
				mapOfArray[r]++
			}
		}

		if len(mapOfArray) == messageSize {
			return index + 1, nil
		}
	}

	return 0, fmt.Errorf("finding marker in '%s'", input)
}

package day06

import (
	"fmt"
)

const packetSize = 4

func FindStartOfPacketMarker(input string) (int, error) {
	foundChars := [packetSize]rune{}

	for index, r := range input {
		foundChars[index%packetSize] = r
		mapOfArray := make(map[rune]int)

		for _, r := range foundChars {
			// Don't map nulls
			if r > 0 {
				mapOfArray[r]++
			}
		}

		if len(mapOfArray) == packetSize {
			return index + 1, nil
		}
	}

	return 0, fmt.Errorf("finding marker in '%s'", input)
}

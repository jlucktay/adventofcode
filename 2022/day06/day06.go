package day06

import "fmt"

const (
	packetSize  = 4
	messageSize = 14
)

func findStartOfMarker(input string, size int) (int, error) {
	foundChars := []rune{}

	for index, r := range input {
		if len(foundChars) < index%size+1 {
			foundChars = append(foundChars, r)
		} else {
			foundChars[index%size] = r
		}

		mapOfArray := make(map[rune]int)

		for _, r := range foundChars {
			// Don't map nulls
			if r > 0 {
				mapOfArray[r]++
			}
		}

		if len(mapOfArray) == size {
			return index + 1, nil
		}
	}

	return 0, fmt.Errorf("finding marker in '%s'", input)
}

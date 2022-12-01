package day01

import (
	"sort"
	"strconv"
	"strings"
)

func ThreeMostCalories(input string) (int, error) {
	xInput := strings.Split(input, "\n")

	currentElfCalories := 0
	mostCalories := []int{0, 0, 0}

	for i := range xInput {
		if len(xInput[i]) == 0 {
			alreadyAddedCalories := false

			for j := range mostCalories {
				if mostCalories[j] == 0 {
					mostCalories[j] = currentElfCalories
					alreadyAddedCalories = true
					break
				}
			}

			for j := range mostCalories {
				if !alreadyAddedCalories && currentElfCalories > mostCalories[j] {
					mostCalories[j] = currentElfCalories
					break
				}
			}

			sort.Ints(mostCalories)

			currentElfCalories = 0

			continue
		}

		lineCalories, err := strconv.ParseInt(xInput[i], 10, 32)
		if err != nil {
			return 0, err
		}

		currentElfCalories += int(lineCalories)
	}

	result := 0

	for k := range mostCalories {
		result += mostCalories[k]
	}

	return result, nil
}

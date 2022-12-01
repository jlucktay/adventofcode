package day01

import (
	"strconv"
	"strings"
)

func MostCalories(input string) (int, error) {
	xInput := strings.Split(input, "\n")

	currentElfCalories, mostCalories := 0, 0

	for i := range xInput {
		if len(xInput[i]) == 0 {
			if currentElfCalories > mostCalories {
				mostCalories = currentElfCalories
			}

			currentElfCalories = 0

			continue
		}

		lineCalories, err := strconv.ParseInt(xInput[i], 10, 32)
		if err != nil {
			return 0, err
		}

		currentElfCalories += int(lineCalories)
	}

	return mostCalories, nil
}

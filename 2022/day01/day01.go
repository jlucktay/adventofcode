package day01

import (
	"sort"
	"strconv"
	"strings"
)

func parseElfCalories(input string) ([]int, error) {
	xInput := strings.Split(input, "\n")
	currentElfCalories := 0
	elfCalories := make([]int, 0)

	for i := range xInput {
		if len(xInput[i]) == 0 && currentElfCalories > 0 {
			elfCalories = append(elfCalories, currentElfCalories)
			currentElfCalories = 0

			continue
		}

		lineCalories, err := strconv.ParseInt(xInput[i], 10, 32)
		if err != nil {
			return nil, err
		}

		currentElfCalories += int(lineCalories)
	}

	sort.Ints(elfCalories)

	return elfCalories, nil
}

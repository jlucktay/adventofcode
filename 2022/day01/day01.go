package day01

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parseElfCalories(input string) ([]int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	currentElfCalories := 0
	elfCalories := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 && currentElfCalories > 0 {
			elfCalories = append(elfCalories, currentElfCalories)
			currentElfCalories = 0

			continue
		}

		lineCalories, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			return nil, err
		}

		currentElfCalories += int(lineCalories)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning input: %v", err)
	}

	if currentElfCalories > 0 {
		elfCalories = append(elfCalories, currentElfCalories)
	}

	sort.Ints(elfCalories)

	return elfCalories, nil
}

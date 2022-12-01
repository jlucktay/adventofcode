package day01

import (
	"errors"
)

func ThreeMostCalories(input string) (int, error) {
	elfCalories, err := parseElfCalories(input)
	if err != nil {
		return 0, err
	}

	if len(elfCalories) < 3 {
		return 0, errors.New("not enough elves")
	}

	accumCalories := 0

	for _, j := range elfCalories[len(elfCalories)-3:] {
		accumCalories += j
	}

	return accumCalories, nil
}

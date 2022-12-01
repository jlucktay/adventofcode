package day01

import (
	"errors"
)

func MostCalories(input string) (int, error) {
	elfCalories, err := parseElfCalories(input)
	if err != nil {
		return 0, err
	}

	if len(elfCalories) == 0 {
		return 0, errors.New("no elves")
	}

	return elfCalories[len(elfCalories)-1], nil
}

// Package day02 for Advent of Code 2016, day 2, part 1.
// https://adventofcode.com/2016/day/2
package day02

import (
	"image"
	"strconv"
)

func Part1(input string) (int, error) {
	inst, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	return inst.bathroomCode(), nil
}

func (i Instructions) bathroomCode() int {
	if len(i) == 0 {
		return 0
	}

	keypad := map[image.Point]int{
		{0, 0}: 1, {1, 0}: 2, {2, 0}: 3,
		{0, 1}: 4, {1, 1}: 5, {2, 1}: 6,
		{0, 2}: 7, {1, 2}: 8, {2, 2}: 9,
	}

	keypadBounds := image.Rect(0, 0, 3, 3)

	result := ""

	// Start from the '5' in the centre.
	currentPosition := image.Pt(1, 1)

	for _, di := range i {
		for _, dir := range di {
			switch dir {
			case Up:
				if currentPosition.Add(image.Pt(0, -1)).In(keypadBounds) {
					currentPosition.Y -= 1
				}

			case Right:
				if currentPosition.Add(image.Pt(1, 0)).In(keypadBounds) {
					currentPosition.X += 1
				}

			case Down:
				if currentPosition.Add(image.Pt(0, 1)).In(keypadBounds) {
					currentPosition.Y += 1
				}

			case Left:
				if currentPosition.Add(image.Pt(-1, 0)).In(keypadBounds) {
					currentPosition.X -= 1
				}

			default:
				panic("unknown direction: " + dir.String())
			}
		}

		result += strconv.Itoa(keypad[currentPosition])
	}

	convertedResult, err := strconv.Atoi(result)
	if err != nil {
		panic("converting '" + result + "': " + err.Error())
	}

	return convertedResult
}

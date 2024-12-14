// Package day02 for Advent of Code 2016, day 2, part 1.
// https://adventofcode.com/2016/day/2
package day02

import "image"

func Part1(input string) (string, error) {
	inst, err := parseInput(input)
	if err != nil {
		return "", err
	}

	/*
		1 2 3
		4 5 6
		7 8 9
	*/

	part1Keypad := map[image.Point]rune{
		{0, 0}: '1', {1, 0}: '2', {2, 0}: '3',
		{0, 1}: '4', {1, 1}: '5', {2, 1}: '6',
		{0, 2}: '7', {1, 2}: '8', {2, 2}: '9',
	}

	return inst.bathroomCode(part1Keypad), nil
}

func (i Instructions) bathroomCode(keypad map[image.Point]rune) string {
	if len(i) == 0 {
		return ""
	}

	currPos := image.Pt(-1, -1)

	// Find the '5' to start from.
	for position, padDigit := range keypad {
		if padDigit == '5' {
			currPos = position
		}
	}

	if currPos.Eq(image.Pt(-1, -1)) {
		panic("could not find '5' on keypad to start from")
	}

	result := ""

	for _, di := range i {
		for _, dir := range di {
			switch dir {
			case Up:
				if _, valid := keypad[currPos.Add(move[Up])]; valid {
					currPos = currPos.Add(move[Up])
				}

			case Right:
				if _, valid := keypad[currPos.Add(move[Right])]; valid {
					currPos = currPos.Add(move[Right])
				}

			case Down:
				if _, valid := keypad[currPos.Add(move[Down])]; valid {
					currPos = currPos.Add(move[Down])
				}

			case Left:
				if _, valid := keypad[currPos.Add(move[Left])]; valid {
					currPos = currPos.Add(move[Left])
				}

			default:
				panic("unknown direction: " + dir.String())
			}
		}

		result += string(keypad[currPos])
	}

	return result
}

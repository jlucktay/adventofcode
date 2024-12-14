// Package day02 for Advent of Code 2016, day 2, part 2.
// https://adventofcode.com/2016/day/2
package day02

import "image"

func Part2(input string) (string, error) {
	inst, err := parseInput(input)
	if err != nil {
		return "", err
	}

	/*
				1
			2 3 4
		5 6 7 8 9
			A B C
				D
	*/

	part2Keypad := map[image.Point]rune{
		{2, 0}: '1',
		{1, 1}: '2', {2, 1}: '3', {3, 1}: '4',
		{0, 2}: '5', {1, 2}: '6', {2, 2}: '7', {3, 2}: '8', {4, 2}: '9',
		{1, 3}: 'A', {2, 3}: 'B', {3, 3}: 'C',
		{2, 4}: 'D',
	}

	return inst.bathroomCode(part2Keypad), nil
}

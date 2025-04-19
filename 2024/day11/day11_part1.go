// Package day11 for Advent of Code 2024, day 11, part 1.
// https://adventofcode.com/2024/day/11
package day11

import (
	"fmt"
	"strconv"
)

func Part1(input string) (int, error) {
	stones, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for range 25 {
		stones = stones.blink()
	}

	result := 0

	for _, count := range stones {
		result += count
	}

	return result, nil
}

func (s Stones) blink() Stones {
	newStones := make(Stones)

	for stone, count := range s {
		for _, b := range stone.blink() {
			newStones[b] += count
		}
	}

	return newStones
}

func (s Stone) blink() []Stone {
	if s == 0 {
		return []Stone{1}
	} else if strStrone := fmt.Sprintf("%d", s); len(strStrone)%2 == 0 {
		leftHalf := strStrone[:len(strStrone)/2]
		rightHalf := strStrone[len(strStrone)/2:]

		leftStone, err := strconv.ParseUint(leftHalf, 10, 64)
		if err != nil {
			panic("couldn't parse left half of " + strStrone)
		}

		rightStone, err := strconv.ParseUint(rightHalf, 10, 64)
		if err != nil {
			panic("couldn't parse right half of " + strStrone)
		}

		return []Stone{Stone(leftStone), Stone(rightStone)}
	} else {
		return []Stone{s * 2024}
	}
}

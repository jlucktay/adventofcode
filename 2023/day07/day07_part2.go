// Package main for Advent of Code 2023, day 7, part 2
// https://adventofcode.com/2023/day/7
package main

func Part2(inputLines []string) (int, error) {
	return parseBids(inputLines, true)
}

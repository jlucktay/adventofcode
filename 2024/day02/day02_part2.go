// Package day02 for Advent of Code 2024, day 2, part 2.
// https://adventofcode.com/2024/day/2
package day02

import (
	"slices"

	"go.jlucktay.dev/adventofcode/crunchy"
)

func Part2(input string) (int, error) {
	reports, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

	for _, report := range reports {
		if (reportStrictlyAsc(report) || reportStrictlyDesc(report)) &&
			reportGapsBetween1And3(report) {

			result++

			continue
		}

		for index := range report {
			indexDeleted := slices.Delete(slices.Clone(report), index, index+1)

			if (reportStrictlyAsc(indexDeleted) || reportStrictlyDesc(indexDeleted)) &&
				reportGapsBetween1And3(indexDeleted) {

				result++

				break
			}
		}
	}

	return result, nil
}

func reportAscOrDesc(report []int) bool {
	if slices.IsSorted(report) {
		return true
	}

	clonedReport := slices.Clone(report)

	slices.Reverse(clonedReport)

	return slices.IsSorted(clonedReport)
}

func reportStrictlyAsc(report []int) bool {
	for index := range report {
		if index >= len(report)-1 {
			break
		}

		if report[index] >= report[index+1] {
			return false
		}
	}

	return true
}

func reportStrictlyDesc(report []int) bool {
	for index := range report {
		if index >= len(report)-1 {
			break
		}

		if report[index] <= report[index+1] {
			return false
		}
	}

	return true
}

func reportHasDuplicates(report []int) bool {
	clone := slices.Clone(report)

	compactedClone := slices.Compact(clone)

	return len(clone)-len(compactedClone) > 0
}

func reportGapsBetween1And3(report []int) bool {
	for index := range report {
		if index >= len(report)-1 {
			break
		}

		if crunchy.AbsDiff(report[index], report[index+1]) >= 1 && crunchy.AbsDiff(report[index], report[index+1]) <= 3 {
			continue
		}

		return false
	}

	return true
}

func reportRemoveFirstUnsafe(report []int) []int {
	clonedReport := slices.Clone(report)

	for index := range clonedReport {
		if index >= len(report)-2 {
			break
		}

		if (clonedReport[index] <= clonedReport[index+1] && clonedReport[index+1] >= clonedReport[index+2]) ||
			(clonedReport[index] >= clonedReport[index+1] && clonedReport[index+1] <= clonedReport[index+2]) {

			firstUnsafeRemoved := slices.Delete(clonedReport, index+1, index+2)

			if reportHasDuplicates(firstUnsafeRemoved) {
				return nil
			}

			return firstUnsafeRemoved
		}
	}

	return nil
}

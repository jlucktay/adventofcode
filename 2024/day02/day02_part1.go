package day02

import (
	"slices"

	"go.jlucktay.dev/adventofcode/crunchy"
)

func SafeReports(input string) (int, error) {
	reports, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result := 0

NextReport:
	for index := range reports {
		report := slices.Clone(reports[index])

		if slices.IsSorted(report) {
			// ascending
		} else {
			if slices.Reverse(report); slices.IsSorted(report) {
				// descending

				slices.Reverse(report)
			} else {
				continue
			}
		}

		lineIndex := 0

		for lineIndex+1 < len(report) {
			v1, v2 := report[lineIndex], report[lineIndex+1]

			if diff := crunchy.AbsDiff(v1, v2); diff > 3 {
				continue NextReport
			}

			if v1 == v2 {
				continue NextReport
			}

			lineIndex++
		}

		result++
	}

	return result, nil
}

// Package day08 for Advent of Code 2024, day 8, part 2.
// https://adventofcode.com/2024/day/8
package day08

func Part2(input string) (int, error) {
	ag, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	ag.plotAntinodes(true)

	return len(ag.thingDirectory[Antinode]), nil
}

func (ag AntennaGrid) inBounds(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if len(ag.theGrid) > 0 && len(ag.theGrid[0]) > 0 {
		return x < len(ag.theGrid[0]) && y < len(ag.theGrid)
	}

	return false
}

func (ag *AntennaGrid) plotAlongLine(gpt GridPointThing, primary [2]int, secondaries [][2]int) {
	if len(secondaries) == 0 {
		return
	}

	for _, secondary := range secondaries {
		xDiff, yDiff := secondary[0]-primary[0], secondary[1]-primary[1]

		xPlot, yPlot := primary[0], primary[1]

		for ag.inBounds(xPlot, yPlot) {
			ag.plotAntinode(xPlot, yPlot)

			xPlot += xDiff
			yPlot += yDiff
		}

		xPlot, yPlot = primary[0], primary[1]

		for ag.inBounds(xPlot, yPlot) {
			ag.plotAntinode(xPlot, yPlot)

			xPlot -= xDiff
			yPlot -= yDiff
		}
	}

	ag.plotAlongLine(gpt, secondaries[0], secondaries[1:])
}

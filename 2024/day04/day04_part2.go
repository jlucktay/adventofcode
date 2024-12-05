// Package day04 for Advent of Code 2024, day 4, part 2.
// https://adventofcode.com/2024/day/4
package day04

func Part2(input string) (int, error) {
	ws, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	return searchWordX(ws), nil
}

func searchX(grid wordSearch, row, col int) bool {
	var m, n int

	// m is horizontal, increasing to the right.
	m = len(grid)
	if m > 0 {
		// n is vertical, increasing going down.
		n = len(grid[0])
	}

	// If the given coordinates do not match the central letter, return false.
	if grid[row][col] != 'A' {
		return false
	}

	// If we're on an edge, return false.
	if col+1 >= m || col-1 < 0 || row+1 >= n || row-1 < 0 {
		return false
	}

	topLeftToBottomRight := string([]rune{grid[row-1][col-1], 'A', grid[row+1][col+1]})
	bottomLeftToTopRight := string([]rune{grid[row+1][col-1], 'A', grid[row-1][col+1]})

	if (topLeftToBottomRight == "MAS" || topLeftToBottomRight == "SAM") &&
		(bottomLeftToTopRight == "MAS" || bottomLeftToTopRight == "SAM") {

		return true
	}

	return false
}

// searchWordX calls searchX for each coordinate.
func searchWordX(grid wordSearch) int {
	var m, n int

	m = len(grid)
	if m > 0 {
		n = len(grid[0])
	}

	result := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if searchX(grid, i, j) {
				result++
			}
		}
	}

	return result
}

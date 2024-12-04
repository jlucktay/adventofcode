// Package day04 for Advent of Code 2024, day 4, part 1.
// https://adventofcode.com/2024/day/4
package day04

func Part1(input string) (int, error) {
	ws, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	return searchWord(ws, "XMAS"), nil
}

func search2D(grid wordSearch, row, col int, word string) int {
	var m, n int

	m = len(grid)
	if m > 0 {
		n = len(grid[0])
	}

	foundCount := 0

	// If the given coordinate do not match the first letter, return false.
	if grid[row][col] != []rune(word)[0] {
		return foundCount
	}

	// x and y are used to set the direction in which to search.
	x := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	y := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	// This loop searches in all 8 directions, one by one.
	// It returns true if any one direction contains the word.
	for dir := 0; dir < 8; dir++ {
		// Initialise the starting point for the current direction.
		var k int
		currX := row + x[dir]
		currY := col + y[dir]

		// The first character was already checked above, so just look for the rest.
		for k = 1; k < len(word); k++ {
			// Break if we're out of bounds.
			if currX >= m || currX < 0 || currY >= n || currY < 0 {
				break
			}

			// Break if the current character doesn't match.
			if grid[currX][currY] != []rune(word)[k] {
				break
			}

			// Move along one in the same direction.
			currX += x[dir]
			currY += y[dir]
		}

		// If all characters matched along the way, this will line up.
		if k == len(word) {
			foundCount++
		}
	}

	return foundCount
}

// searchWord calls search2D for each coordinate.
func searchWord(grid wordSearch, word string) int {
	var m, n int

	m = len(grid)
	if m > 0 {
		n = len(grid[0])
	}

	result := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result += search2D(grid, i, j, word)
		}
	}

	return result
}

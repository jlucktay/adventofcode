// Package main for Advent of Code 2023, day 10, part 2
// https://adventofcode.com/2023/day/10
package main

// ShoelacePicks combines the [shoelace formula] with [Pick's theorem] to calculate the area inside the pipe loop.
//
// [shoelace formula]: https://en.wikipedia.org/wiki/Shoelace_formula
// [Pick's theorem]: https://en.wikipedia.org/wiki/Pick%27s_theorem
func (p Pipes) ShoelacePicks() int {
	pts := make([]*Tile, 0)

	if p.start == nil {
		return 0
	}

	pts = append(pts, p.start)

	for pipe := p.start.next; pipe != p.start; pipe = pipe.next {
		pts = append(pts, pipe)
	}

	// https://en.wikipedia.org/wiki/Shoelace_formula
	area := 0

	p0 := pts[len(pts)-1]

	for _, p1 := range pts {
		area += p0.x*p1.y - p0.y*p1.x
		p0 = p1
	}

	if area < 0 {
		area = -area
	}

	area /= 2

	// Transpose for 'i': https://en.wikipedia.org/wiki/Pick's_theorem#Formula
	return area - len(pts)/2 + 1
}

func Part2(inputLines []string) (int, error) {
	grid, err := parseInput(inputLines)
	if err != nil {
		return 0, err
	}

	pipes, err := parsePipes(grid)
	if err != nil {
		return 0, err
	}

	return pipes.ShoelacePicks(), nil
}

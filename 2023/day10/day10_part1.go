// Package main for Advent of Code 2023, day 10, part 1
// https://adventofcode.com/2023/day/10
package main

import (
	"errors"
	"fmt"
)

// Pipes represents the "one large, continuous loop" parsed from the input.
type Pipes struct {
	start *Tile
}

func (p Pipes) SolveForPart1() int {
	length := 1

	if p.start == nil {
		return 0
	}

	if p.start.next == nil {
		return 1
	}

	current := p.start.next

	for current != p.start {
		current = current.next
		length++
	}

	return length / 2
}

// Tile is one section of pipe, connected in two directions.
type Tile struct {
	// next Tile/Pipe in the loop.
	next *Tile

	// prev Tile/Pipe in the loop.
	prev *Tile

	// kind of Pipe on this Tile.
	kind Kind

	// Coordinates of this Tile in the grid.
	x, y int
}

type Kind rune

const (
	Vertical         Kind = '|'
	Horizontal       Kind = '-'
	NorthAndEast     Kind = 'L'
	NorthAndWest     Kind = 'J'
	SouthAndWest     Kind = '7'
	SouthAndEast     Kind = 'F'
	Ground           Kind = '.'
	StartingPosition Kind = 'S'
)

type Direction string

const (
	North Direction = "north"
	East  Direction = "east"
	South Direction = "south"
	West  Direction = "west"
)

var cardinals = []Direction{North, East, South, West}

// Grid is for parsing the puzzle input map.
type Grid [][]*Tile

func parseLine(inputLine string, lineNumber int) ([]*Tile, error) {
	result := make([]*Tile, 0)

	for ix, r := range inputLine {
		switch r {
		case '|':
			result = append(result, &Tile{kind: Vertical, x: ix, y: lineNumber})
		case '-':
			result = append(result, &Tile{kind: Horizontal, x: ix, y: lineNumber})
		case 'L':
			result = append(result, &Tile{kind: NorthAndEast, x: ix, y: lineNumber})
		case 'J':
			result = append(result, &Tile{kind: NorthAndWest, x: ix, y: lineNumber})
		case '7':
			result = append(result, &Tile{kind: SouthAndWest, x: ix, y: lineNumber})
		case 'F':
			result = append(result, &Tile{kind: SouthAndEast, x: ix, y: lineNumber})
		case '.':
			result = append(result, &Tile{kind: Ground, x: ix, y: lineNumber})
		case 'S':
			result = append(result, &Tile{kind: StartingPosition, x: ix, y: lineNumber})
		default:
			return nil, fmt.Errorf("unknown rune '%s'", string(r))
		}
	}

	return result, nil
}

func parseInput(inputLines []string) (Grid, error) {
	grid := Grid{}

	for il, line := range inputLines {
		parsed, err := parseLine(line, il)
		if err != nil {
			return nil, err
		}

		if len(parsed) > 0 {
			grid = append(grid, parsed)
		}
	}

	return grid, nil
}

func parsePipes(grid Grid) (Pipes, error) {
	pipes := Pipes{}

	if len(grid) == 0 {
		return pipes, nil
	}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x].kind == StartingPosition {
				pipes.start = grid[y][x]

				break
			}
		}

		if pipes.start != nil {
			break
		}
	}

	if pipes.start == nil {
		return Pipes{}, errors.New("could not find starting position in grid")
	}

	// Parse the pipes in the loop.
	connsFromStart, err := pipes.start.LookAroundForConnections(grid)
	if err != nil {
		return Pipes{}, err
	}

	current := connsFromStart[0]
	current.prev = pipes.start

	pipes.start.next = current

	for current != pipes.start {
		// Sanity check, just in case.
		if current.prev.next != current {
			return Pipes{}, errors.New("current.prev.next is not current")
		}

		connections, err := current.LookAroundForConnections(grid)
		if err != nil {
			return Pipes{}, err
		}

		if current.prev != connections[0] && current.prev != connections[1] {
			return pipes, errors.New("neither possible connection is the previous one")
		}

		switch current.prev {
		case connections[0]:
			current.next = connections[1]
		case connections[1]:
			current.next = connections[0]
		default:
			return pipes, errors.New("neither connection was the previous one")
		}

		tmp := current
		current = current.next
		current.prev = tmp
	}

	return pipes, nil
}

// LookAroundForConnections will return a slice of Tiles pointers of length 2, or an error.
func (t *Tile) LookAroundForConnections(grid Grid) ([]*Tile, error) {
	// Discover (up to) 4 neighbours.
	neighbours, err := t.GetNeighbours(grid)
	if err != nil {
		return nil, err
	}

	// Exactly 2 neighbours should connect to this Tile.
	result := make([]*Tile, 0)

	for _, direction := range cardinals {
		neighbour, neighbourExists := neighbours[direction]

		if neighbourExists {
			if t.ConnectsTo(neighbour) {
				result = append(result, neighbour)
			}
		}
	}

	if len(result) != 2 {
		return nil, fmt.Errorf("found %d connections instead of 2", len(result))
	}

	return result, nil
}

func (k Kind) ConnectsTo(d Direction, l Kind) bool {
	switch k {
	case Vertical:
		switch d {
		case North:
			return l == StartingPosition || l == Vertical || l == SouthAndEast || l == SouthAndWest

		case East:
			return false

		case South:
			return l == StartingPosition || l == Vertical || l == NorthAndEast || l == NorthAndWest

		case West:
			return false
		}

	case Horizontal:
		switch d {
		case North:
			return false

		case East:
			return l == StartingPosition || l == Horizontal || l == NorthAndWest || l == SouthAndWest

		case South:
			return false

		case West:
			return l == StartingPosition || l == Horizontal || l == NorthAndEast || l == SouthAndEast
		}

	case NorthAndEast:
		switch d {
		case North:
			return l == StartingPosition || l == Vertical || l == SouthAndEast || l == SouthAndWest

		case East:
			return l == StartingPosition || l == Horizontal || l == NorthAndWest || l == SouthAndWest

		case South:
			return false

		case West:
			return false
		}

	case NorthAndWest:
		switch d {
		case North:
			return l == StartingPosition || l == Vertical || l == SouthAndEast || l == SouthAndWest

		case East:
			return false

		case South:
			return false

		case West:
			return l == StartingPosition || l == Horizontal || l == NorthAndEast || l == SouthAndEast
		}

	case SouthAndWest:
		switch d {
		case North:
			return false

		case East:
			return false

		case South:
			return l == StartingPosition || l == Vertical || l == NorthAndEast || l == NorthAndWest

		case West:
			return l == StartingPosition || l == Horizontal || l == NorthAndEast || l == SouthAndEast
		}

	case SouthAndEast:
		switch d {
		case North:
			return false

		case East:
			return l == StartingPosition || l == Horizontal || l == NorthAndWest || l == SouthAndWest

		case South:
			return l == StartingPosition || l == Vertical || l == NorthAndEast || l == NorthAndWest

		case West:
			return false
		}

	case Ground:
		return false

	case StartingPosition:
		switch d {
		case North:
			return l == Vertical || l == SouthAndEast || l == SouthAndWest

		case East:
			return l == Horizontal || l == NorthAndWest || l == SouthAndWest

		case South:
			return l == Vertical || l == NorthAndEast || l == NorthAndWest

		case West:
			return l == Horizontal || l == NorthAndEast || l == SouthAndEast
		}
	}

	return false
}

func (t *Tile) ConnectsTo(u *Tile) bool {
	switch {
	case t.y-1 == u.y: // connecting up/north
		return t.kind.ConnectsTo(North, u.kind)

	case t.x+1 == u.x: // connecting right/east
		return t.kind.ConnectsTo(East, u.kind)

	case t.y+1 == u.y: // connecting down/south
		return t.kind.ConnectsTo(South, u.kind)

	case t.x-1 == u.x: // connecting left/west
		return t.kind.ConnectsTo(West, u.kind)

	default:
		return false
	}
}

func (t *Tile) GetNeighbours(g Grid) (map[Direction]*Tile, error) {
	result := make(map[Direction]*Tile)

	if t.x-1 >= 0 {
		// Does not fall off the left/west edge.
		result[West] = g[t.y][t.x-1]
	}

	if t.x+1 < len(g[t.y]) {
		// Does not fall off the right/east edge.
		result[East] = g[t.y][t.x+1]
	}

	if t.y-1 >= 0 {
		// Does not fall off the top/north edge.
		result[North] = g[t.y-1][t.x]
	}

	if t.y+1 < len(g) {
		// Does not fall off the bottom/south edge.
		result[South] = g[t.y+1][t.x]
	}

	return result, nil
}

func Part1(inputLines []string) (int, error) {
	grid, err := parseInput(inputLines)
	if err != nil {
		return 0, err
	}

	pipes, err := parsePipes(grid)
	if err != nil {
		return 0, err
	}

	return pipes.SolveForPart1(), nil
}

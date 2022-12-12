package day12

import (
	"fmt"
	"strings"
)

type Tile struct {
	// Kind is the kind of tile, potentially affecting movement.
	Kind rune

	// X and Y are the coordinates of the tile.
	X, Y int

	// W is a reference to the World that the tile is a part of.
	W World
}

func (t *Tile) neighbours() []*Tile {
	neighbours := []*Tile{}

	for _, offset := range [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
		if n := t.W.Tile(t.X+offset[0], t.Y+offset[1]); n != nil {
			from := t.Kind
			switch from {
			case 'S':
				from = 'a'
			case 'E':
				from = 'z'
			}

			to := n.Kind
			switch to {
			case 'S':
				to = 'a'
			case 'E':
				to = 'z'
			}

			if from <= to+1 {
				neighbours = append(neighbours, n)
			}
		}
	}

	return neighbours
}

// World is a two dimensional map of Tiles.
type World map[int]map[int]*Tile

// Tile gets the tile at the given coordinates in the world.
func (w World) Tile(x, y int) *Tile {
	if w[x] == nil {
		return nil
	}

	return w[x][y]
}

// SetTile sets a tile at the given coordinates in the world.
func (w World) SetTile(t *Tile, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*Tile{}
	}

	w[x][y] = t
	t.X = x
	t.Y = y
	t.W = w
}

// FirstOfKind gets the first tile on the board of a kind.
// Used to get the from and to tiles as there should only be one of each.
func (w World) FirstOfKind(kind rune) *Tile {
	for _, row := range w {
		for _, t := range row {
			if t.Kind == kind {
				return t
			}
		}
	}

	return nil
}

// From gets the from tile from the world.
func (w World) From() *Tile {
	return w.FirstOfKind('S')
}

// To gets the to tile from the world.
func (w World) To() *Tile {
	return w.FirstOfKind('E')
}

// RenderPath renders a path on top of a world.
func (w World) RenderPath(path []*Tile) string {
	width := len(w)
	if width == 0 {
		return ""
	}

	height := len(w[0])
	pathLocs := map[string]bool{}

	for _, p := range path {
		pathLocs[fmt.Sprintf("%d,%d", p.X, p.Y)] = true
	}

	rows := make([]string, height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			t := w.Tile(x, y)
			r := ' '

			if pathLocs[fmt.Sprintf("%d,%d", x, y)] && t.Kind != 'E' {
				r = '*'
			} else if t != nil {
				r = t.Kind
			}

			rows[y] += string(r)
		}
	}

	return strings.Join(rows, "\n") + "\n"
}

// ParseWorld parses a textual representation of a world into a world map.
func ParseWorld(input string) World {
	w := World{}

	for y, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, raw := range row {
			w.SetTile(&Tile{Kind: raw}, x, y)
		}
	}

	return w
}

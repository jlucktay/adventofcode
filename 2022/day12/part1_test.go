package day12_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day12"
)

const INPUT = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

func TestPart1(t *testing.T) {
	is := is.New(t)

	actual, err := day12.Part1(INPUT)
	is.NoErr(err)
	is.Equal(31, actual)
}

func TestParseWorld(t *testing.T) {
	is := is.New(t)

	w := day12.ParseWorld(INPUT)
	is.Equal(INPUT, w.RenderPath([]*day12.Tile{}))
}

func TestTabulaRasa(t *testing.T) {
	is := is.New(t)

	w := day12.ParseWorld(INPUT)
	is.Equal('S', w.FirstOfKind('S').Kind)
	is.Equal('S', w.From().Kind)
	is.Equal('S', w.Tile(0, 0).Kind)
	is.Equal('b', w.Tile(1, 1).Kind)
	is.Equal('c', w.Tile(2, 2).Kind)
	is.Equal('t', w.Tile(3, 3).Kind)
	is.Equal('f', w.Tile(4, 4).Kind)
}

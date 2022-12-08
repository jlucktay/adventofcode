package day08

import (
	"testing"

	"github.com/matryer/is"
)

const INPUT = `30373
25512
65332
33549
35390
`

func TestNewTallTreeGrid(t *testing.T) {
	is := is.New(t)

	actual, err := NewTallTreeGrid(INPUT)
	is.NoErr(err)
	is.Equal(
		[][]int{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		},
		actual.trees,
	)
}

func TestLookDownAt(t *testing.T) {
	is := is.New(t)

	actual, err := NewTallTreeGrid(INPUT)
	is.NoErr(err)

	is.Equal(actual.lookDownAt(1, 1), true)
	is.Equal(actual.lookDownAt(2, 1), true)
	is.Equal(actual.lookDownAt(3, 1), false)

	is.Equal(actual.lookDownAt(1, 2), false)
	is.Equal(actual.lookDownAt(2, 2), false)
	is.Equal(actual.lookDownAt(3, 2), false)

	is.Equal(actual.lookDownAt(1, 3), false)
	is.Equal(actual.lookDownAt(2, 3), false)
	is.Equal(actual.lookDownAt(3, 3), false)
}

func TestLookLeftAt(t *testing.T) {
	is := is.New(t)

	actual, err := NewTallTreeGrid(INPUT)
	is.NoErr(err)

	is.Equal(actual.lookLeftAt(1, 1), false)
	is.Equal(actual.lookLeftAt(2, 1), true)
	is.Equal(actual.lookLeftAt(3, 1), false)

	is.Equal(actual.lookLeftAt(1, 2), true)
	is.Equal(actual.lookLeftAt(2, 2), false)
	is.Equal(actual.lookLeftAt(3, 2), true)

	is.Equal(actual.lookLeftAt(1, 3), false)
	is.Equal(actual.lookLeftAt(2, 3), false)
	is.Equal(actual.lookLeftAt(3, 3), false)
}

func TestLookUpAt(t *testing.T) {
	is := is.New(t)

	actual, err := NewTallTreeGrid(INPUT)
	is.NoErr(err)

	is.Equal(actual.lookUpAt(1, 1), false)
	is.Equal(actual.lookUpAt(2, 1), false)
	is.Equal(actual.lookUpAt(3, 1), false)

	is.Equal(actual.lookUpAt(1, 2), false)
	is.Equal(actual.lookUpAt(2, 2), false)
	is.Equal(actual.lookUpAt(3, 2), false)

	is.Equal(actual.lookUpAt(1, 3), false)
	is.Equal(actual.lookUpAt(2, 3), true)
	is.Equal(actual.lookUpAt(3, 3), false)
}

func TestLookRightAt(t *testing.T) {
	is := is.New(t)

	actual, err := NewTallTreeGrid(INPUT)
	is.NoErr(err)

	is.Equal(actual.lookRightAt(1, 1), true)
	is.Equal(actual.lookRightAt(2, 1), false)
	is.Equal(actual.lookRightAt(3, 1), false)

	is.Equal(actual.lookRightAt(1, 2), false)
	is.Equal(actual.lookRightAt(2, 2), false)
	is.Equal(actual.lookRightAt(3, 2), false)

	is.Equal(actual.lookRightAt(1, 3), false)
	is.Equal(actual.lookRightAt(2, 3), true)
	is.Equal(actual.lookRightAt(3, 3), false)
}

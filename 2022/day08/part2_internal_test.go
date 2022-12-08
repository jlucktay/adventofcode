package day08

import (
	"testing"

	"github.com/matryer/is"
)

func TestLookUpFrom(t *testing.T) {
	is := is.New(t)

	actual, err := NewTallTreeGrid(INPUT)
	is.NoErr(err)

	is.Equal(actual.lookUpFrom(0, 0), 0)
	is.Equal(actual.lookUpFrom(1, 0), 0)
	is.Equal(actual.lookUpFrom(2, 0), 0)
	is.Equal(actual.lookUpFrom(3, 0), 0)
	is.Equal(actual.lookUpFrom(4, 0), 0)

	is.Equal(actual.lookUpFrom(0, 1), 1)
	is.Equal(actual.lookUpFrom(1, 1), 1)
	is.Equal(actual.lookUpFrom(2, 1), 1)
	is.Equal(actual.lookUpFrom(3, 1), 1)
	is.Equal(actual.lookUpFrom(4, 1), 1)

	is.Equal(actual.lookUpFrom(0, 2), 2)
	is.Equal(actual.lookUpFrom(1, 2), 1)
	is.Equal(actual.lookUpFrom(2, 2), 1)
	is.Equal(actual.lookUpFrom(3, 2), 2)
	is.Equal(actual.lookUpFrom(4, 2), 1)

	is.Equal(actual.lookUpFrom(0, 3), 1)
	is.Equal(actual.lookUpFrom(1, 3), 1)
	is.Equal(actual.lookUpFrom(2, 3), 2)
	is.Equal(actual.lookUpFrom(3, 3), 3)
	is.Equal(actual.lookUpFrom(4, 3), 3)

	is.Equal(actual.lookUpFrom(0, 4), 1)
	is.Equal(actual.lookUpFrom(1, 4), 2)
	is.Equal(actual.lookUpFrom(2, 4), 1)
	is.Equal(actual.lookUpFrom(3, 4), 4)
	is.Equal(actual.lookUpFrom(4, 4), 1)
}

func TestLookRightFrom(t *testing.T) {
	is := is.New(t)

	actual, err := NewTallTreeGrid(INPUT)
	is.NoErr(err)

	is.Equal(actual.lookRightFrom(0, 0), 2)
	is.Equal(actual.lookRightFrom(1, 0), 1)
	is.Equal(actual.lookRightFrom(2, 0), 1)
	is.Equal(actual.lookRightFrom(3, 0), 1)
	is.Equal(actual.lookRightFrom(4, 0), 0)

	is.Equal(actual.lookRightFrom(0, 1), 1)
	is.Equal(actual.lookRightFrom(1, 1), 1)
	is.Equal(actual.lookRightFrom(2, 1), 2)
	is.Equal(actual.lookRightFrom(3, 1), 1)
	is.Equal(actual.lookRightFrom(4, 1), 0)

	is.Equal(actual.lookRightFrom(0, 2), 4)
	is.Equal(actual.lookRightFrom(1, 2), 3)
	is.Equal(actual.lookRightFrom(2, 2), 1)
	is.Equal(actual.lookRightFrom(3, 2), 1)
	is.Equal(actual.lookRightFrom(4, 2), 0)

	is.Equal(actual.lookRightFrom(0, 3), 1)
	is.Equal(actual.lookRightFrom(1, 3), 1)
	is.Equal(actual.lookRightFrom(2, 3), 2)
	is.Equal(actual.lookRightFrom(3, 3), 1)
	is.Equal(actual.lookRightFrom(4, 3), 0)

	is.Equal(actual.lookRightFrom(0, 4), 1)
	is.Equal(actual.lookRightFrom(1, 4), 2)
	is.Equal(actual.lookRightFrom(2, 4), 1)
	is.Equal(actual.lookRightFrom(3, 4), 1)
	is.Equal(actual.lookRightFrom(4, 4), 0)
}

func TestLookDownFrom(t *testing.T) {
	is := is.New(t)

	actual, err := NewTallTreeGrid(INPUT)
	is.NoErr(err)

	is.Equal(actual.lookDownFrom(0, 0), 2)
	is.Equal(actual.lookDownFrom(1, 0), 1)
	is.Equal(actual.lookDownFrom(2, 0), 1)
	is.Equal(actual.lookDownFrom(3, 0), 4)
	is.Equal(actual.lookDownFrom(4, 0), 3)

	is.Equal(actual.lookDownFrom(0, 1), 1)
	is.Equal(actual.lookDownFrom(1, 1), 1)
	is.Equal(actual.lookDownFrom(2, 1), 2)
	is.Equal(actual.lookDownFrom(3, 1), 1)
	is.Equal(actual.lookDownFrom(4, 1), 1)

	is.Equal(actual.lookDownFrom(0, 2), 2)
	is.Equal(actual.lookDownFrom(1, 2), 2)
	is.Equal(actual.lookDownFrom(2, 2), 1)
	is.Equal(actual.lookDownFrom(3, 2), 1)
	is.Equal(actual.lookDownFrom(4, 2), 1)

	is.Equal(actual.lookDownFrom(0, 3), 1)
	is.Equal(actual.lookDownFrom(1, 3), 1)
	is.Equal(actual.lookDownFrom(2, 3), 1)
	is.Equal(actual.lookDownFrom(3, 3), 1)
	is.Equal(actual.lookDownFrom(4, 3), 1)

	is.Equal(actual.lookDownFrom(0, 4), 0)
	is.Equal(actual.lookDownFrom(1, 4), 0)
	is.Equal(actual.lookDownFrom(2, 4), 0)
	is.Equal(actual.lookDownFrom(3, 4), 0)
	is.Equal(actual.lookDownFrom(4, 4), 0)
}

func TestLookLeftFrom(t *testing.T) {
	is := is.New(t)

	actual, err := NewTallTreeGrid(INPUT)
	is.NoErr(err)

	is.Equal(actual.lookLeftFrom(0, 0), 0)
	is.Equal(actual.lookLeftFrom(1, 0), 1)
	is.Equal(actual.lookLeftFrom(2, 0), 2)
	is.Equal(actual.lookLeftFrom(3, 0), 3)
	is.Equal(actual.lookLeftFrom(4, 0), 1)

	is.Equal(actual.lookLeftFrom(0, 1), 0)
	is.Equal(actual.lookLeftFrom(1, 1), 1)
	is.Equal(actual.lookLeftFrom(2, 1), 1)
	is.Equal(actual.lookLeftFrom(3, 1), 1)
	is.Equal(actual.lookLeftFrom(4, 1), 2)

	is.Equal(actual.lookLeftFrom(0, 2), 0)
	is.Equal(actual.lookLeftFrom(1, 2), 1)
	is.Equal(actual.lookLeftFrom(2, 2), 1)
	is.Equal(actual.lookLeftFrom(3, 2), 1)
	is.Equal(actual.lookLeftFrom(4, 2), 1)

	is.Equal(actual.lookLeftFrom(0, 3), 0)
	is.Equal(actual.lookLeftFrom(1, 3), 1)
	is.Equal(actual.lookLeftFrom(2, 3), 2)
	is.Equal(actual.lookLeftFrom(3, 3), 1)
	is.Equal(actual.lookLeftFrom(4, 3), 4)

	is.Equal(actual.lookLeftFrom(0, 4), 0)
	is.Equal(actual.lookLeftFrom(1, 4), 1)
	is.Equal(actual.lookLeftFrom(2, 4), 1)
	is.Equal(actual.lookLeftFrom(3, 4), 3)
	is.Equal(actual.lookLeftFrom(4, 4), 1)
}

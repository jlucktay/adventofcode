package day07_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day07"
)

func TestFindDirToDelete(t *testing.T) {
	is := is.New(t)

	actual, err := day07.FindDirToDelete(day07.INPUT)
	is.NoErr(err)
	is.Equal(24933642, actual)
}

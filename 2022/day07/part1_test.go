package day07_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day07"
)

const PARSED_INPUT = `- / (dir)
  - a (dir)
    - e (dir)
      - i (file, size=584)
    - f (file, size=29116)
    - g (file, size=2557)
    - h.lst (file, size=62596)
  - b.txt (file, size=14848514)
  - c.dat (file, size=8504156)
  - d (dir)
    - d.ext (file, size=5626152)
    - d.log (file, size=8033020)
    - j (file, size=4060174)
    - k (file, size=7214296)
`

func TestParseFileSystem(t *testing.T) {
	is := is.New(t)

	actual, err := day07.ParseFileSystem(day07.INPUT)
	is.NoErr(err)
	is.Equal(PARSED_INPUT, actual.String())
}

func TestFindDirsUpTo100K(t *testing.T) {
	is := is.New(t)

	actual, err := day07.FindDirsUpTo100K(day07.INPUT)
	is.NoErr(err)
	is.Equal(95437, actual)
}

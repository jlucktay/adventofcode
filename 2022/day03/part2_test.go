package day03_test

import (
	"bufio"
	"errors"
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day03"
)

func TestRucksackGroupPriority(t *testing.T) {
	is := is.New(t)

	actual, err := day03.RucksackGroupPriority(INPUT)
	is.NoErr(err)
	is.Equal(70, actual)
}

func TestNewScanner(t *testing.T) {
	is := is.New(t)

	testInputs := map[string]struct {
		eof bool
		str []byte
	}{
		"empty": {
			eof: false,
			str: []byte(""),
		},
		"one line only": {
			eof: false,
			str: []byte(`one
`),
		},
		"stuff": {
			eof: false,
			str: []byte(`one
two
three
`),
		},
		"more stuff": {
			eof: true,
			str: []byte(`one
two
three
`),
		},
		"lots of lines - EOF true": {
			eof: true,
			str: []byte(`one
two
three
four
five
six
`),
		},
		"lots of lines - EOF false": {
			eof: false,
			str: []byte(`one
two
three
four
five
six
`),
		},
	}

	for desc, testInput := range testInputs {
		desc, testInput := desc, testInput
		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			is := is.New(t)

			_, _, err := day03.SplitThreeLines(testInput.str, testInput.eof)

			is.True(errors.Is(bufio.ErrFinalToken, err) || err == nil)
		})
	}
}

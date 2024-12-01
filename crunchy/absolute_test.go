package crunchy_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/crunchy"
)

func TestAbsolute(t *testing.T) {
	is := is.New(t)

	is.Equal(crunchy.AbsoluteDiff(42, 27), 15)
	is.Equal(crunchy.AbsoluteDiff(27, 42), 15)
}

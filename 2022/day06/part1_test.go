package day06_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day06"
)

func TestFindStartOfPacketMarker(t *testing.T) {
	is := is.New(t)

	for _, input := range INPUT {
		marker, err := day06.FindStartOfPacketMarker(input.datastream)
		is.NoErr(err)
		is.Equal(input.packet, marker)
	}
}

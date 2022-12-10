package day10_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day10"
)

const PRODUCED_IMAGE = `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
`

func TestRenderPixels(t *testing.T) {
	is := is.New(t)

	cc := day10.NewClockCircuit()

	is.NoErr(cc.ParseProgram(LARGER_PROGRAM))

	is.Equal(PRODUCED_IMAGE, cc.RenderPixels())
}

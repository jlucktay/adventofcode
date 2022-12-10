package day10_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2022/day10"
)

/*
Consider the following small program:

noop
addx 3
addx -5

Execution of this program proceeds as follows:

    At the start of the first cycle, the noop instruction begins execution. During the first cycle, X is 1. After the first cycle, the noop instruction finishes execution, doing nothing.
    At the start of the second cycle, the addx 3 instruction begins execution. During the second cycle, X is still 1.
    During the third cycle, X is still 1. After the third cycle, the addx 3 instruction finishes execution, setting X to 4.
    At the start of the fourth cycle, the addx -5 instruction begins execution. During the fourth cycle, X is still 4.
    During the fifth cycle, X is still 4. After the fifth cycle, the addx -5 instruction finishes execution, setting X to -1.
*/

func TestSmallProgram(t *testing.T) {
	is := is.New(t)

	cc := day10.NewClockCircuit()
	cc.QueueProgram(day10.InstNoop, 0)
	cc.QueueProgram(day10.InstAddX, 3)
	cc.QueueProgram(day10.InstAddX, -5)

	// At the start of the first cycle, the noop instruction begins execution.
	is.Equal(1, cc.Value())

	// During the first cycle, X is 1.
	// After the first cycle, the noop instruction finishes execution, doing nothing.
	cc.Tick()
	is.Equal(1, cc.Value())

	// At the start of the second cycle, the addx 3 instruction begins execution.
	// During the second cycle, X is still 1.
	cc.Tick()
	is.Equal(1, cc.Value())

	// During the third cycle, X is still 1.
	// After the third cycle, the addx 3 instruction finishes execution, setting X to 4.
	cc.Tick()
	is.Equal(4, cc.Value())

	// At the start of the fourth cycle, the addx -5 instruction begins execution.
	// During the fourth cycle, X is still 4.
	cc.Tick()
	is.Equal(4, cc.Value())

	// During the fifth cycle, X is still 4.
	// After the fifth cycle, the addx -5 instruction finishes execution, setting X to -1.
	cc.Tick()
	is.Equal(-1, cc.Value())

	tape := cc.Tape()
	is.Equal(1, tape[1])
	is.Equal(1, tape[2])
	is.Equal(4, tape[3])
	is.Equal(4, tape[4])
	is.Equal(-1, tape[5])
}

const LARGER_PROGRAM = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
`

func TestConsiderLargerProgram(t *testing.T) {
	is := is.New(t)

	cc := day10.NewClockCircuit()
	is.NoErr(cc.ParseProgram(LARGER_PROGRAM))

	tape := cc.Tape()
	is.Equal(420, tape[19]*20)    // 21
	is.Equal(1140, tape[59]*60)   // 19
	is.Equal(1800, tape[99]*100)  // 18
	is.Equal(2940, tape[139]*140) // 21
	is.Equal(2880, tape[179]*180) // 16
	is.Equal(3960, tape[219]*220) // expect 18 * 220 = 3960
}

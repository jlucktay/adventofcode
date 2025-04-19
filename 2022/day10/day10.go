package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ClockCircuit struct {
	cycle, index uint
	register     int
	program      []programLine
	busyAdding   bool
}

type programLine struct {
	i         instruction
	parameter int
}

type instruction string

const (
	InstNoop instruction = "noop"
	InstAddX instruction = "addx"
)

func NewClockCircuit() ClockCircuit {
	return ClockCircuit{
		cycle:      0,
		index:      0,
		register:   1,
		program:    []programLine{},
		busyAdding: false,
	}
}

func (cc *ClockCircuit) QueueProgram(inst instruction, param int) {
	cc.program = append(cc.program, programLine{inst, param})
}

func (cc *ClockCircuit) Tick() {
	if cc.index >= uint(len(cc.program)) {
		fmt.Fprintf(os.Stderr, "index %d out of bounds", cc.index)

		return
	}

	if cc.busyAdding {
		cc.register += cc.program[cc.index].parameter
		cc.index++
		cc.cycle++
		cc.busyAdding = false

		return
	}

	switch cc.program[cc.index].i {
	case InstNoop:
	case InstAddX:
		cc.busyAdding = true
		cc.cycle++

		return
	default:
		return
	}

	cc.index++
	cc.cycle++
}

func (cc *ClockCircuit) Cycle() uint {
	return cc.cycle
}

func (cc *ClockCircuit) Value() int {
	return cc.register
}

func (cc *ClockCircuit) ParseProgram(input string) error {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		xLine := strings.Split(line, " ")

		if len(xLine) < 1 || len(xLine) > 2 {
			return fmt.Errorf("parsing '%s'", line)
		}

		switch instruction(xLine[0]) {
		case InstNoop:
			cc.QueueProgram(InstNoop, 0)
		case InstAddX:
			param, err := strconv.ParseInt(xLine[1], 10, 32)
			if err != nil {
				return fmt.Errorf("parsing integer: %w", err)
			}

			cc.QueueProgram(InstAddX, int(param))
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning input: %w", err)
	}

	return nil
}

func (cc *ClockCircuit) Tape() []int {
	result := []int{1}

	for _, pl := range cc.program {
		result = append(result, result[len(result)-1])

		switch pl.i {
		case InstNoop:
		case InstAddX:
			result = append(result, result[len(result)-1]+pl.parameter)
		}
	}

	return result
}

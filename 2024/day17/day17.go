// Package day17 for Advent of Code 2024, day 17.
// https://adventofcode.com/2024/day/17
package day17

import (
	"fmt"
	"strconv"
	"strings"
)

type register = int

type ChronospatialComputer struct {
	A, B, C            register
	rawProgram         string
	Program            []int
	instructionPointer int
	rawOutput          []int
}

func (cc *ChronospatialComputer) comboOperand(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand

	case 4:
		return cc.A

	case 5:
		return cc.B

	case 6:
		return cc.C

	case 7:
		panic("combo operand 7 is reserved and will not appear in valid programs.")

	default:
		panic("invalid combo operand '" + strconv.Itoa(operand) + "'")
	}
}

func (cc *ChronospatialComputer) operand() int {
	if cc.instructionPointer >= len(cc.Program) {
		panic("instruction pointer is past the length of the program")
	}

	opcode := cc.Program[cc.instructionPointer]
	literalOperand := cc.Program[cc.instructionPointer+1]

	switch opcode {
	case 0, 2, 5, 6, 7:
		return cc.comboOperand(literalOperand)

	case 1, 3:
		return literalOperand

	case 4:
		// no-op; opcode 4 ignores operand
		return 0

	default:
		panic("unknown opcode: " + strconv.Itoa(opcode))
	}
}

func (cc *ChronospatialComputer) Process() bool {
	if cc.instructionPointer >= len(cc.Program) {
		return false
	}

	switch cc.Program[cc.instructionPointer] {
	case 0:
		cc.ADV(cc.operand())

	case 1:
		cc.BXL(cc.operand())

	case 2:
		cc.BST(cc.operand())

	case 3:
		cc.JNZ(cc.operand())

	case 4:
		cc.BXC(cc.operand())

	case 5:
		cc.OUT(cc.operand())

	case 6:
		cc.BDV(cc.operand())

	case 7:
		cc.CDV(cc.operand())
	}

	return true
}

// ADV (opcode 0) performs division. The numerator is the value in the A register. The denominator is found by raising
// 2 to the power of the instruction's combo operand. (So, an operand of 2 would divide A by 4 (2^2); an operand of 5
// would divide A by 2^B.) The result of the division operation is truncated to an integer and then written to the A
// register.
func (cc *ChronospatialComputer) ADV(operand int) {
	cc.A /= (1 << operand)
	cc.instructionPointer += 2
}

// BXL (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand, then stores the
// result in register B.
func (cc *ChronospatialComputer) BXL(operand int) {
	cc.B ^= operand
	cc.instructionPointer += 2
}

// BST (opcode 2) calculates the value of its combo operand modulo 8 (thereby keeping only its lowest 3 bits), then
// writes that value to the B register.
func (cc *ChronospatialComputer) BST(operand int) {
	cc.B = operand % 8
	cc.instructionPointer += 2
}

// JNZ (opcode 3) does nothing if the A register is 0. However, if the A register is not zero, it jumps by setting the
// instruction pointer to the value of its literal operand; if this instruction jumps, the instruction pointer is not
// increased by 2 after this instruction.
func (cc *ChronospatialComputer) JNZ(operand int) {
	if cc.A == 0 {
		cc.instructionPointer += 2
		return
	}

	cc.instructionPointer = operand
}

// BXC (opcode 4) calculates the bitwise XOR of register B and register C, then stores the result in register B. (For
// legacy reasons, this instruction reads an operand but ignores it.)
func (cc *ChronospatialComputer) BXC(operand int) {
	cc.B ^= cc.C
	cc.instructionPointer += 2
}

// OUT (opcode 5) calculates the value of its combo operand modulo 8, then outputs that value. (If a program outputs
// multiple values, they are separated by commas.)
func (cc *ChronospatialComputer) OUT(operand int) {
	cc.rawOutput = append(cc.rawOutput, operand%8)
	cc.instructionPointer += 2
}

// BDV (opcode 6) works exactly like the adv instruction except that the result is stored in the B register. (The
// numerator is still read from the A register.)
func (cc *ChronospatialComputer) BDV(operand int) {
	cc.B = cc.A / (1 << operand)
	cc.instructionPointer += 2
}

// CDV (opcode 7) works exactly like the adv instruction except that the result is stored in the C register. (The
// numerator is still read from the A register.)
func (cc *ChronospatialComputer) CDV(operand int) {
	cc.C = cc.A / (1 << operand)
	cc.instructionPointer += 2
}

func (cc *ChronospatialComputer) Output() string {
	converted := make([]string, 0)

	for _, rawOut := range cc.rawOutput {
		converted = append(converted, strconv.Itoa(rawOut))
	}

	return strings.Join(converted, ",")
}

func parseInput(input string) (ChronospatialComputer, error) {
	if len(input) == 0 {
		return ChronospatialComputer{}, nil
	}

	xInput := strings.Split(input, "\n\n")
	if len(xInput) != 2 {
		return ChronospatialComputer{},
			fmt.Errorf("input '%s' split into %d fields instead of 2 (registers and program)", input, len(xInput))
	}

	registers := xInput[0]
	program := xInput[1]

	xRegisters := strings.Split(registers, "\n")
	if len(xRegisters) != 3 {
		return ChronospatialComputer{},
			fmt.Errorf("registers '%s' split into %d fields instead of 3 (A,B,C)", registers, len(xRegisters))
	}

	result := ChronospatialComputer{
		rawOutput: make([]int, 0),
	}

	for _, regLine := range xRegisters {
		var reg string
		var value int

		n, err := fmt.Sscanf(regLine, "Register %1s: %d", &reg, &value)
		if n != 2 || err != nil {
			return ChronospatialComputer{}, fmt.Errorf("scanning register '%s' parsed %d items: %w", regLine, n, err)
		}

		switch reg {
		case "A":
			result.A = value
		case "B":
			result.B = value
		case "C":
			result.C = value
		}
	}

	if !strings.HasPrefix(program, "Program: ") {
		return ChronospatialComputer{}, fmt.Errorf("program '%s' does not start with 'Program: '", program)
	}

	xProgram := strings.Split(program, " ")
	if len(xProgram) != 2 {
		return ChronospatialComputer{}, fmt.Errorf("program '%s' split into %d fields instead of 2", program, len(xProgram))
	}

	result.rawProgram = strings.TrimSpace(xProgram[1])

	xPrograms := strings.Split(result.rawProgram, ",")
	if len(xPrograms)%2 != 0 {
		return ChronospatialComputer{},
			fmt.Errorf("splitting raw programs '%s' did not give an even number of tokens", xProgram[1])
	}

	for _, rawProgram := range xPrograms {
		parsedProgram, err := strconv.Atoi(rawProgram)
		if err != nil {
			return ChronospatialComputer{}, fmt.Errorf("converting string '%s' to int: %w", rawProgram, err)
		}

		result.Program = append(result.Program, parsedProgram)
	}

	return result, nil
}

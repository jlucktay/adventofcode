package day11

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Monkey struct {
	inventory       []int
	op              action
	testDivFactor   int
	ifTrue, ifFalse int
	thrownCounter   int
}

type action struct {
	op     operation
	factor int
}

type operation rune

const (
	opMultiply operation = '*'
	opAdd      operation = '+'
)

func ParseMonkey(input string) (*Monkey, error) {
	m := &Monkey{
		inventory: []int{},
		op:        action{},
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	var err error

	// Monkey 0:
	_ = scanner.Scan()

	// Starting items: 79, 98
	if !scanner.Scan() {
		return nil, errors.New("scanning starting items line")
	}
	siLine := scanner.Text()
	siLine = strings.TrimPrefix(siLine, "  Starting items: ")
	for _, item := range strings.Split(siLine, ", ") {
		itemNumber, err := strconv.Atoi(item)
		if err != nil {
			return nil, fmt.Errorf("starting items '%#v': %w", siLine, err)
		}

		m.inventory = append(m.inventory, itemNumber)
	}

	// Operation: new = old * 19
	if !scanner.Scan() {
		return nil, errors.New("scanning operation line")
	}
	opLine := scanner.Text()
	opLine = strings.TrimPrefix(opLine, "  Operation: new = old ")
	ops := strings.Split(opLine, " ")
	r, _ := utf8.DecodeRuneInString(ops[0])
	m.op.op = operation(r)

	if strings.TrimSpace(ops[1]) == "old" {
		// Special case
		m.op.factor = -1
	} else {
		m.op.factor, err = strconv.Atoi(ops[1])
		if err != nil {
			return nil, fmt.Errorf("op factor '%#v': %w", ops[1], err)
		}
	}

	// Test: divisible by 23
	if !scanner.Scan() {
		return nil, errors.New("scanning test line")
	}
	testLine := scanner.Text()
	testLine = strings.TrimPrefix(testLine, "  Test: divisible by ")
	m.testDivFactor, err = strconv.Atoi(testLine)
	if err != nil {
		return nil, fmt.Errorf("test '%#v': %w", testLine, err)
	}

	// If true: throw to monkey 2
	if !scanner.Scan() {
		return nil, errors.New("scanning if true line")
	}
	ifTrueLine := scanner.Text()
	ifTrueLine = strings.TrimPrefix(ifTrueLine, "    If true: throw to monkey ")
	m.ifTrue, err = strconv.Atoi(ifTrueLine)
	if err != nil {
		return nil, fmt.Errorf("if true '%#v': %w", ifTrueLine, err)
	}

	// If false: throw to monkey 3
	if !scanner.Scan() {
		return nil, errors.New("scanning if false line")
	}
	ifFalseLine := scanner.Text()
	ifFalseLine = strings.TrimPrefix(ifFalseLine, "    If false: throw to monkey ")
	m.ifFalse, err = strconv.Atoi(ifFalseLine)
	if err != nil {
		return nil, fmt.Errorf("if false '%#v': %w", ifFalseLine, err)
	}

	return m, nil
}

func (m *Monkey) Turn() (map[int][]int, error) {
	throwTo := map[int][]int{}

	for _, item := range m.inventory {
		// Apply monkey's operation.
		factor := m.op.factor

		if m.op.factor == -1 {
			factor = item
		}

		switch m.op.op {
		case opMultiply:
			item *= factor
		case opAdd:
			item += factor
		default:
			return nil, fmt.Errorf("unknown op '%s'", string(m.op.op))
		}

		// Divide by three and round down.
		item /= 3

		// Test worry level.
		if item%m.testDivFactor == 0 {
			if _, exists := throwTo[m.ifTrue]; !exists {
				throwTo[m.ifTrue] = make([]int, 0)
			}

			throwTo[m.ifTrue] = append(throwTo[m.ifTrue], item)
		} else {
			if _, exists := throwTo[m.ifFalse]; !exists {
				throwTo[m.ifFalse] = make([]int, 0)
			}

			throwTo[m.ifFalse] = append(throwTo[m.ifFalse], item)
		}
	}

	// Inventory has all been thrown away now.
	m.thrownCounter += len(m.inventory)
	m.inventory = []int{}

	return throwTo, nil
}

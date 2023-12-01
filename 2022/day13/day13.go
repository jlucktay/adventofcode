package day13

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func parse(input string) ([][2]PacketData, error) {
	result := [][2]PacketData{}

	pairs := strings.Split(input, "\n\n")

	for pairIndex := range pairs {
		xLine := strings.Split(pairs[pairIndex], "\n")

		left, _, err := parseList(xLine[0], 0)
		if err != nil {
			return nil, err
		}

		right, _, err := parseList(xLine[1], 0)
		if err != nil {
			return nil, err
		}

		result = append(result, [2]PacketData{left, right})
	}

	return result, nil
}

func parseList(input string, indent int) (PacketData, int, error) {
	pl := &PacketList{}
	index := 0

	var intBuffer strings.Builder

	for index < len(input) {
		r, size := utf8.DecodeRuneInString(input[index:])

		if r == ']' || r == ',' {
			ibs := intBuffer.String()

			if len(ibs) > 0 {
				parseThis, err := strconv.Atoi(ibs)
				if err != nil {
					return nil, 0, fmt.Errorf("parsing '%s': %w", ibs, err)
				}

				intBuffer.Reset()
				*pl = append(*pl, PacketInteger(parseThis))
			}

			if r == ']' {
				return pl, index, nil
			}
		}

		if r == '[' {
			substring := input[index+1:]

			innerPL, iplSize, err := parseList(input[index+1:], indent+1)
			if err != nil {
				return nil, 0, err
			}

			ipl, ok := innerPL.(*PacketList)
			if !ok {
				return nil, 0, fmt.Errorf("asserting PacketList type on '%#v'", innerPL)
			}

			*pl = append(*pl, *ipl...)

			index += iplSize + 1
			continue
		}

		if unicode.IsDigit(r) {
			intBuffer.WriteRune(r)
		}

		index += size
	}

	return nil, 0, errors.New("reached end of input string without matching all contained lists")
}

/*
Packet data consists of lists and integers.
Each list starts with `[`, ends with `]`, and contains zero or more comma-separated values
(either integers or other lists).
Each packet is always a list and appears on its own line.
*/

type PacketData interface {
	LowerThan(PacketData) bool
}

type PacketList []PacketData

func (pl *PacketList) LowerThan(other PacketData) bool {
	return false
}

type PacketInteger int

func (pi PacketInteger) LowerThan(other PacketData) bool {
	pl := PacketList{pi}
	return pl.LowerThan(other)
}

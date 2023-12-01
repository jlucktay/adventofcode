package day13

import (
	"errors"
	"fmt"
	"log"
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

// parseList takes an input line (or part of one) and parses through it until it finds the end of the list, calling
// itself recursively as necessary if and when it finds the start of another inner list.
// The parsed list is returned, along with the number of runes parsed.
// If there are any errors along the way, the error is returned instead.
func parseList(input string, indent int) (PacketData, int, error) {
	// create a blank list
	// iterate through each rune
	// is this '['?
	//   get the substring from here to the end and call recursively
	//   advance to rune after the next ']' that was found
	// is this a digit?
	//   add to a buffer to strconv later
	// is this a comma?
	//   strconv the buffer and add result to current list

	pl := &PacketList{}
	index := 0

	var intBuffer strings.Builder

	//
	log.Printf("% *snew parseList call; input '%s'",
		indent, "", input)

	defer func() {
		log.Printf("% *sparseList call done; input '%s'",
			indent, "", input)
	}()
	//

	for index < len(input) {
		r, size := utf8.DecodeRuneInString(input[index:])

		//
		log.Printf("% *sgot rune '%s' at index %d (of %d total length) of input '%s'",
			indent, "", string(r), index, len(input), input)
		//

		if r == ']' || r == ',' {
			ibs := intBuffer.String()

			if len(ibs) > 0 {
				parseThis, err := strconv.Atoi(ibs)
				if err != nil {
					return nil, 0, fmt.Errorf("parsing '%s': %w", ibs, err)
				}

				intBuffer.Reset()
				*pl = append(*pl, PacketInteger(parseThis))

				//
				log.Printf("% *sconverted buffer and added %d to list",
					indent, "", parseThis)
				//
			}

			if r == ']' {
				//
				log.Printf("% *sfound a closing bracket; converted buffer, returning '%#v'",
					indent, "", *pl)
				//

				return pl, index, nil
			}

			if r == ',' {
				//
				log.Printf("% *sfound a comma; converted buffer",
					indent, "")
				//
			}
		}

		if r == '[' {
			substring := input[index+1:]

			//
			log.Printf("% *snew list inside this list: '%s' - calling recursively",
				indent, "", substring)
			//

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
			//
			log.Printf("% *s'%s' is a digit, buffering for conversion later",
				indent, "", string(r))
			//

			intBuffer.WriteRune(r)
		}

		//
		log.Printf("% *sparsed list: '%#v'",
			indent, "", *pl)
		//

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

/*
When comparing two values, the first value is called left and the second value is called right. Then:

  - If both values are integers, the lower integer should come first. If the left integer is lower than the right
    integer, the inputs are in the right order. If the left integer is higher than the right integer, the inputs are
    not in the right order. Otherwise, the inputs are the same integer; continue checking the next part of the input.

  - If both values are lists, compare the first value of each list, then the second value, and so on. If the left list
    runs out of items first, the inputs are in the right order. If the right list runs out of items first, the inputs
    are not in the right order. If the lists are the same length and no comparison makes a decision about the order,
    continue checking the next part of the input.

  - If exactly one value is an integer, convert the integer to a list which contains that integer as its only value,
    then retry the comparison. For example, if comparing [0,0,0] and 2, convert the right value to [2] (a list
    containing 2); the result is then found by instead comparing [0,0,0] and [2].
*/

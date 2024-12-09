// Package day09 for Advent of Code 2024, day 9, part 1.
// https://adventofcode.com/2024/day/9
package day09

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Disk []int64

func (d Disk) String() string {
	sb := strings.Builder{}

	for _, block := range d {
		var printLetter rune

		if block == -1 {
			printLetter = '.'
		} else if block > 10 {
			printLetter = 'x'
		} else {
			pl, _ := utf8.DecodeRuneInString(fmt.Sprintf("%d", block))
			printLetter = pl
		}

		sb.WriteRune(printLetter)
	}

	return sb.String()
}

func (d Disk) firstFree() int {
	for index, block := range d {
		if block == -1 {
			return index
		}
	}

	return -1
}

func (d Disk) lastFile() int {
	for i := len(d) - 1; i >= 0; i-- {
		if d[i] != -1 {
			return i
		}
	}

	return -1
}

func (d Disk) moveSingleBlock(from, to int) {
	fileID := d[from]
	d[from] = -1
	d[to] = fileID
}

func (d Disk) compact() {
	if len(d) == 0 {
		return
	}

	for d.firstFree() != d.lastFile()+1 {
		d.moveSingleBlock(d.lastFile(), d.firstFree())
	}
}

func (d Disk) checksum(part2 bool) int64 {
	result := int64(0)

	for i := int64(0); i < int64(len(d)); i++ {
		if d[i] == -1 {
			if part2 {
				continue
			} else {
				break
			}
		}

		result += i * d[i]
	}

	return result
}

func Part1(input string) (int64, error) {
	d, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	d.compact()

	return d.checksum(false), nil
}

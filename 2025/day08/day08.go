// Package day08 for Advent of Code 2025, day 8.
// https://adventofcode.com/2025/day/8
package day08

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"log/slog"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Playground struct {
	boxes     []*junctionBox
	distances []distancePair
	circuits  map[int][]*junctionBox
}

type distancePair struct {
	left, right *junctionBox
	distance    int
}

func (p Playground) String() string {
	sb := strings.Builder{}

	for _, box := range p.boxes {
		sb.WriteString(box.String())
		sb.WriteRune('\n')
	}

	for _, distPair := range p.distances {
		sb.WriteString(fmt.Sprintf("\t%s<->%s: %d\n", distPair.left, distPair.right, distPair.distance))
	}

	sb.WriteString("Circuits:\n")

	for circuitNumber, boxesInCircuit := range p.circuits {
		sb.WriteString(fmt.Sprintf("[%d]: ", circuitNumber))

		for cIdx := range len(boxesInCircuit) {
			sb.WriteString(boxesInCircuit[cIdx].String())

			if cIdx < len(boxesInCircuit)-1 {
				sb.WriteRune(',')
			} else {
				sb.WriteRune('\n')
			}
		}
	}

	return sb.String()
}

// stopAfter making this many connections.
// This value changes between the example and the real input.
// Thank you: https://www.reddit.com/r/adventofcode/comments/1ph6aul/2025_day_8_part_1_outofband_extra_information/
func (p *Playground) stopAfter() int {
	if len(p.boxes) == 20 {
		return 10
	} else if len(p.boxes) == 1000 {
		return 1000
	} else {
		return 0
	}
}

func (p *Playground) largestThreeCircuits() int {
	largest := []int{0, 0, 0}

	slog.Debug("looking for three largest circuits", slog.Any("circuits", p.circuits))

	for _, circuit := range p.circuits {
		slog.Debug("considering next circuit", slog.Any("circuit", circuit), slog.Int("size", len(circuit)))

		if len(circuit) > largest[2] {
			largest[2] = len(circuit)
		}

		slices.Sort(largest)
		slices.Reverse(largest)

		slog.Debug("after considering most recent circuit", slog.Any("largest", largest))
	}

	slog.Debug("found three largest circuits", slog.Any("largest", largest))

	return largest[0] * largest[1] * largest[2]
}

type junctionBox struct {
	coords  [3]int
	circuit int
}

func (jb *junctionBox) String() string {
	sb := strings.Builder{}

	sb.WriteRune('(')
	sb.WriteString(fmt.Sprintf("%5d", jb.coords[jbX]))
	sb.WriteRune(',')
	sb.WriteString(fmt.Sprintf("%5d", jb.coords[jbY]))
	sb.WriteRune(',')
	sb.WriteString(fmt.Sprintf("%5d", jb.coords[jbZ]))
	sb.WriteRune(')')

	return sb.String()
}

const (
	jbX = iota
	jbY
	jbZ
)

func (jb *junctionBox) EuclideanDistance(target *junctionBox) int {
	xDelta := (jb.coords[jbX] - target.coords[jbX]) * (jb.coords[jbX] - target.coords[jbX])
	yDelta := (jb.coords[jbY] - target.coords[jbY]) * (jb.coords[jbY] - target.coords[jbY])
	zDelta := (jb.coords[jbZ] - target.coords[jbZ]) * (jb.coords[jbZ] - target.coords[jbZ])

	return int(math.Sqrt(float64(xDelta) + float64(yDelta) + float64(zDelta)))
}

func (p *Playground) calculateDistances() {
	for outerIdx := range p.boxes {
	NextInner:
		for innerIdx := outerIdx + 1; innerIdx < len(p.boxes); innerIdx++ {
			for _, dp := range p.distances {
				if (dp.left == p.boxes[outerIdx] && dp.right == p.boxes[innerIdx]) || (dp.left == p.boxes[innerIdx] && dp.right == p.boxes[outerIdx]) {
					slog.Debug("already know distance between these two boxes; skipping", slog.String("outer", p.boxes[outerIdx].String()), slog.String("inner", p.boxes[innerIdx].String()), slog.Int("distance", dp.distance))

					break NextInner
				}
			}

			slog.Debug("distance between these two boxes is not yet known; calculating", slog.String("outer", p.boxes[outerIdx].String()), slog.String("inner", p.boxes[innerIdx].String()))

			newDistCalc := p.boxes[outerIdx].EuclideanDistance(p.boxes[innerIdx])

			newDistPair := distancePair{
				left:     p.boxes[outerIdx],
				right:    p.boxes[innerIdx],
				distance: newDistCalc,
			}

			slog.Debug("distance between these two boxes is now calculated", slog.Any("new", newDistPair))

			p.distances = append(p.distances, newDistPair)
		}
	}

	slices.SortStableFunc(p.distances, func(a, b distancePair) int {
		return cmp.Compare(a.distance, b.distance)
	})
}

func (p *Playground) connectClosestPairs() {
	nextCircuit := 0

	for sa := range p.stopAfter() {
		if p.distances[sa].left.circuit == -1 && p.distances[sa].right.circuit == -1 {
			p.distances[sa].left.circuit = nextCircuit
			p.distances[sa].right.circuit = nextCircuit

			nextCircuit++
		} else if p.distances[sa].left.circuit == -1 && p.distances[sa].right.circuit != -1 {
			p.distances[sa].left.circuit = p.distances[sa].right.circuit
		} else if p.distances[sa].right.circuit == -1 && p.distances[sa].left.circuit != -1 {
			p.distances[sa].right.circuit = p.distances[sa].left.circuit
		} else {
			// Both are already part of existing circuits, so collapse to the lower of the two.
			lowestCircuit := int(math.Min(float64(p.distances[sa].left.circuit), float64(p.distances[sa].right.circuit)))
			highestCircuit := int(math.Max(float64(p.distances[sa].left.circuit), float64(p.distances[sa].right.circuit)))

			for circuitUpdateIdx := range p.stopAfter() {
				if p.distances[circuitUpdateIdx].left.circuit == highestCircuit {
					p.distances[circuitUpdateIdx].left.circuit = lowestCircuit
				}

				if p.distances[circuitUpdateIdx].right.circuit == highestCircuit {
					p.distances[circuitUpdateIdx].right.circuit = lowestCircuit
				}
			}
		}
	}

	for boxIdx := range p.boxes {
		if p.boxes[boxIdx].circuit == -1 {
			continue
		}

		p.circuits[p.boxes[boxIdx].circuit] = append(p.circuits[p.boxes[boxIdx].circuit], p.boxes[boxIdx])
	}
}

func parseInput(input string) (Playground, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := Playground{
		boxes:     make([]*junctionBox, 0),
		distances: make([]distancePair, 0),
		circuits:  make(map[int][]*junctionBox),
	}

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), ",")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		if len(xLine) != 3 {
			return Playground{}, fmt.Errorf("unexpected line format '%+v'", xLine)
		}

		newBox := &junctionBox{
			coords:  [3]int{},
			circuit: -1,
		}

		for index, rawCoord := range xLine {
			coord, err := strconv.Atoi(rawCoord)
			if err != nil {
				return Playground{}, fmt.Errorf("converting co-ordinate '%s': %w", rawCoord, err)
			}

			newBox.coords[index] = coord
		}

		result.boxes = append(result.boxes, newBox)
	}

	return result, nil
}

// Package day03 for Advent of Code 2025, day 3.
// https://adventofcode.com/2025/day/3
package day03

import (
	"bufio"
	"bytes"
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"strings"
)

type BatteryBank []int

type BatteryBanks []BatteryBank

func (bb BatteryBank) largestJoltage() int {
	highestTotal := 0

	for i := 0; i < len(bb)-1; i++ {
		left := bb[i]

		for j := i + 1; j < len(bb); j++ {
			right := bb[j]

			if candidate := left*10 + right; candidate > highestTotal {
				highestTotal = candidate
			}
		}
	}

	return highestTotal
}

// overcomeStaticFriction was made possible by this write-up: https://www.geeksforgeeks.org/dsa/largest-number-possible-after-removal-of-k-digits/
func (bb BatteryBank) overcomeStaticFriction() (int64, error) {
	sb := strings.Builder{}

	for _, b := range bb {
		sb.WriteString(fmt.Sprintf("%d", b))
	}

	sbs := sb.String()

	n, ok := new(big.Int).SetString(sbs, 10)
	if !ok {
		return 0, fmt.Errorf("setting string '%s' to big.Int", sbs)
	}

	k := len(bb) - 12

	if k <= 0 {
		return 0, fmt.Errorf("battery bank '%+v' is not big enough (%d) to turn on 12 batteries", bb, len(bb))
	}

	// Generate the largest number after removal of the least K digits one by one
	for range k {
		ans := big.NewInt(0)
		i := big.NewInt(1)

		// Remove the least digit after every iteration
		for new(big.Int).Div(n, i).Cmp(big.NewInt(0)) == 1 {
			// Store the numbers formed after removing every digit once
			iTimesTen := new(big.Int).Mul(i, big.NewInt(10))
			nOverI := new(big.Int).Div(n, iTimesTen)
			leftOfPlus := new(big.Int).Mul(nOverI, i)

			rightOfPlus := new(big.Int).Mod(n, i)

			temp := new(big.Int).Add(leftOfPlus, rightOfPlus)

			i.Mul(i, big.NewInt(10))

			// Compare and store the maximum
			if temp.Cmp(ans) == 1 {
				ans = temp
			}
		}

		// Store the largest number remaining
		n = ans
	}

	return n.Int64(), nil
}

func parseInput(input string) (BatteryBanks, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := BatteryBanks{}

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), " ")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		if len(xLine) != 1 {
			return nil, fmt.Errorf("unexpected format for line '%+v'", xLine)
		}

		newBank := make(BatteryBank, 0)

		for _, r := range xLine[0] {
			newBattery, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, fmt.Errorf("converting battery value '%v'", r)
			}

			newBank = append(newBank, newBattery)
		}

		result = append(result, newBank)
	}

	return result, nil
}

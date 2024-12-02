package day02

import (
	"bufio"
	"bytes"
	"slices"
	"strconv"
	"strings"
)

type reports = [][]int

func parseInput(input string) (reports, error) {
	buffer := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buffer)

	result := make([][]int, 0)
	row := 0

	for scanner.Scan() {
		xLine := strings.Split(scanner.Text(), " ")
		xLine = slices.DeleteFunc(xLine, func(s string) bool { return s == "" })

		result = append(result, make([]int, 0))

		for index := range xLine {
			num, err := strconv.Atoi(xLine[index])
			if err != nil {
				return nil, err
			}

			result[row] = append(result[row], num)
		}

		row++
	}

	return result, nil
}

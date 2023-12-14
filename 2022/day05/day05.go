package day05

import (
	"fmt"
	"strings"
)

// Big thanks to:
// https://leetcode.com/problems/sliding-window-maximum/discuss/693349/golang-deque-with-explanation

type crateDeque []rune

func (cd *crateDeque) append(element rune) {
	*cd = append(*cd, element)
}

func (cd *crateDeque) prepend(element rune) {
	*cd = append(crateDeque{element}, *cd...)
}

func (cd *crateDeque) getFirst() (rune, bool) {
	if cd.empty() {
		return 0, false
	}

	return (*cd)[0], true
}

func (cd *crateDeque) popFirst() (rune, bool) {
	if cd.empty() {
		return 0, false
	}

	element := (*cd)[0]
	*cd = (*cd)[1:]

	return element, true
}

// func (cd *crateDeque) getLast() (rune, bool) {
// 	if cd.empty() {
// 		return 0, false
// 	}
//
// 	return (*cd)[len(*cd)-1], true
// }

// func (cd *crateDeque) popLast() (rune, bool) {
// 	if cd.empty() {
// 		return 0, false
// 	}
//
// 	index := len(*cd) - 1
// 	element := (*cd)[index]
// 	*cd = (*cd)[:index]
//
// 	return element, true
// }

func (cd *crateDeque) empty() bool {
	return len(*cd) == 0
}

func (cd *crateDeque) String() string {
	sb := strings.Builder{}

	sb.WriteString("[")

	for i := 0; i < len(*cd); i++ {
		sb.WriteString(fmt.Sprintf(" %s", string((*cd)[i])))
	}

	sb.WriteString(" ]")

	return sb.String()
}

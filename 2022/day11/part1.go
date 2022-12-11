package day11

import (
	"errors"
	"sort"
	"strings"
)

func TwentyRoundsOfMonkeyBusiness(input string) (int, error) {
	xInput := strings.Split(input, "\n\n")

	monkeys := []Monkey{}

	for _, monkey := range xInput {
		m, err := ParseMonkey(monkey)
		if err != nil {
			return 0, err
		}

		monkeys = append(monkeys, *m)
	}

	for round := 0; round < 20; round++ {
		for index := range monkeys {
			thrownItems, err := monkeys[index].Turn()
			if err != nil {
				return 0, err
			}

			for recipient, items := range thrownItems {
				monkeys[recipient].inventory = append(monkeys[recipient].inventory, items...)
			}
		}
	}

	throwCounters := []int{}

	for index := range monkeys {
		throwCounters = append(throwCounters, monkeys[index].thrownCounter)
	}

	if len(throwCounters) < 2 {
		return 0, errors.New("less than two monkeys threw things")
	}

	sort.Ints(throwCounters)

	result := throwCounters[len(throwCounters)-1] * throwCounters[len(throwCounters)-2]

	return result, nil
}

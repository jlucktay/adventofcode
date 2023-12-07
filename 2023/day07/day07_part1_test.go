/*
To play Camel Cards, you are given a list of hands and their corresponding bid (your puzzle input). For example:

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483

This example shows five hands; each hand is followed by its bid amount. Each hand wins an amount equal to its bid multiplied by its rank, where the weakest hand gets rank 1, the second-weakest hand gets rank 2, and so on up to the strongest hand. Because there are five hands in this example, the strongest hand will have rank 5 and its bid will be multiplied by 5.

So, the first step is to put the hands in order of strength:

    32T3K is the only one pair and the other hands are all a stronger type, so it gets rank 1.
    KK677 and KTJJT are both two pair. Their first cards both have the same label, but the second card of KK677 is stronger (K vs T), so KTJJT gets rank 2 and KK677 gets rank 3.
    T55J5 and QQQJA are both three of a kind. QQQJA has a stronger first card, so it gets rank 5 and T55J5 gets rank 4.

Now, you can determine the total winnings of this set of hands by adding up the result of multiplying each hand's bid with its rank (765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5). So the total winnings in this example are 6440.
*/

package main

import (
	"strings"
	"testing"

	"github.com/matryer/is"

	_ "embed"
)

//go:embed testdata/camel_cards.txt
var camelCards string

func TestPart1(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   []string
		want int
	}{
		"empty": {[]string{}, 0},
		"input": {
			in:   strings.Split(camelCards, "\n"),
			want: 6440,
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := Part1(testCase.in)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}

func TestOfAKind(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in        Hand
		wantCount int
		wantCard  Card
	}{
		"five": {
			in:        Hand{cards: [5]Card{'A', 'A', 'A', 'A', 'A'}},
			wantCount: 5,
			wantCard:  'A',
		},
		"four": {
			in:        Hand{cards: [5]Card{'A', 'A', '8', 'A', 'A'}},
			wantCount: 4,
			wantCard:  'A',
		},
		"three": {
			in:        Hand{cards: [5]Card{'T', 'T', 'T', '9', '8'}},
			wantCount: 3,
			wantCard:  'T',
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			gotBool, gotCard := testCase.in.xOfAKind(testCase.wantCount, ' ')
			is.True(gotBool)
			is.Equal(gotCard, testCase.wantCard)
		})
	}
}

func TestFullHouse(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   Hand
		want bool
	}{
		"three and two": {
			in:   Hand{cards: [5]Card{'2', '3', '3', '3', '2'}},
			want: true,
		},
		"full house": {
			in:   Hand{cards: [5]Card{'2', '2', '2', '2', '2'}},
			want: false,
		},
		"all different": {
			in:   Hand{cards: [5]Card{'2', '3', '4', '5', '6'}},
			want: false,
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got := testCase.in.fullHouse()
			is.Equal(testCase.want, got)
		})
	}
}

func TestTwoPair(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in                                   Hand
		want                                 bool
		wantFirst, wantSecond, wantRemaining Card
	}{
		"three and two": {
			in:            Hand{cards: [5]Card{'2', '3', '4', '3', '2'}},
			want:          true,
			wantFirst:     '2',
			wantSecond:    '3',
			wantRemaining: '4',
		},

		"full house": {
			in:   Hand{cards: [5]Card{'2', '2', '2', '2', '2'}},
			want: false,
		},

		"all different": {
			in:   Hand{cards: [5]Card{'2', '3', '4', '5', '6'}},
			want: false,
		},

		"KK677": {
			in:            Hand{cards: [5]Card{'K', 'K', '6', '7', '7'}},
			want:          true,
			wantFirst:     'K',
			wantSecond:    '7',
			wantRemaining: '6',
		},

		"KTJJT": {
			in:            Hand{cards: [5]Card{'K', 'T', 'J', 'J', 'T'}},
			want:          true,
			wantFirst:     'J',
			wantSecond:    'T',
			wantRemaining: 'K',
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, gotFirst, gotSecond, gotRemaining := testCase.in.twoPair()
			is.Equal(got, testCase.want)

			if testCase.want {
				is.True(gotFirst != gotSecond)
				is.True(testCase.wantFirst == gotFirst || testCase.wantFirst == gotSecond)
				is.True(testCase.wantSecond == gotFirst || testCase.wantSecond == gotSecond)
				is.Equal(string(testCase.wantRemaining), string(gotRemaining))
			}
		})
	}
}

func TestOnePair(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in                                                       Hand
		want                                                     bool
		wantPair, wantRemaining1, wantRemaining2, wantRemaining3 Card
	}{
		"aces": {
			in:             Hand{cards: [5]Card{'A', '2', '3', 'A', '4'}},
			want:           true,
			wantPair:       'A',
			wantRemaining1: '2',
			wantRemaining2: '3',
			wantRemaining3: '4',
		},

		"full house": {
			in:   Hand{cards: [5]Card{'2', '2', '2', '2', '2'}},
			want: false,
		},

		"all different": {
			in:   Hand{cards: [5]Card{'2', '3', '4', '5', '6'}},
			want: false,
		},

		"KK678": {
			in:             Hand{cards: [5]Card{'K', 'K', '6', '7', '8'}},
			want:           true,
			wantPair:       'K',
			wantRemaining1: '6',
			wantRemaining2: '7',
			wantRemaining3: '8',
		},

		"KTJQT": {
			in:             Hand{cards: [5]Card{'K', 'T', 'J', 'Q', 'T'}},
			want:           true,
			wantPair:       'T',
			wantRemaining1: 'K',
			wantRemaining2: 'J',
			wantRemaining3: 'Q',
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, gotPair, gotRemaining1, gotRemaining2, gotRemaining3 := testCase.in.onePair()
			is.Equal(got, testCase.want)

			if testCase.want {
				is.Equal(gotPair, testCase.wantPair)

				is.True(testCase.wantRemaining1 == gotRemaining1 ||
					testCase.wantRemaining1 == gotRemaining2 ||
					testCase.wantRemaining1 == gotRemaining3)

				is.True(testCase.wantRemaining2 == gotRemaining1 ||
					testCase.wantRemaining2 == gotRemaining2 ||
					testCase.wantRemaining2 == gotRemaining3)

				is.True(testCase.wantRemaining3 == gotRemaining1 ||
					testCase.wantRemaining3 == gotRemaining2 ||
					testCase.wantRemaining3 == gotRemaining3)

				is.True(gotRemaining1 != gotRemaining2)
				is.True(gotRemaining2 != gotRemaining3)
			}
		})
	}
}

func TestHighCard(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in    Hand
		joker bool
		want  Card
	}{
		"aces": {
			in:    Hand{cards: [5]Card{'A', '2', '3', '4', '5'}},
			joker: false,
			want:  'A',
		},

		"full house": {
			in:    Hand{cards: [5]Card{'2', '2', '2', '2', '2'}},
			joker: false,
			want:  Err,
		},

		"all different": {
			in:    Hand{cards: [5]Card{'2', '3', '4', '5', '6'}},
			joker: false,
			want:  '6',
		},

		"KK678": {
			in:    Hand{cards: [5]Card{'K', 'K', '6', '7', '8'}},
			joker: false,
			want:  Err,
		},

		"KQ678": {
			in:    Hand{cards: [5]Card{'K', 'Q', '6', '7', '8'}},
			joker: false,
			want:  'K',
		},

		"KTJQT": {
			in:    Hand{cards: [5]Card{'K', 'T', 'J', 'Q', 'T'}},
			joker: false,
			want:  Err,
		},

		"KTJQT jokers wild": {
			in:    Hand{cards: [5]Card{'K', 'T', 'J', 'Q', 'T'}},
			joker: true,
			want:  Err,
		},

		"KTJQ3 jokers wild": {
			in:    Hand{cards: [5]Card{'K', 'T', 'J', 'Q', '3'}},
			joker: true,
			want:  'K',
		},

		"34JQT": {
			in:    Hand{cards: [5]Card{'3', '4', 'J', 'Q', 'T'}},
			joker: false,
			want:  'Q',
		},

		"34JQT jokers wild": {
			in:    Hand{cards: [5]Card{'3', '4', 'J', 'Q', 'T'}},
			joker: true,
			want:  'Q',
		},

		"34JQ5 jokers wild": {
			in:    Hand{cards: [5]Card{'3', '4', 'J', 'Q', '5'}},
			joker: true,
			want:  'Q',
		},

		"34J75 jokers wild": {
			in:    Hand{cards: [5]Card{'3', '4', 'J', '7', '5'}},
			joker: true,
			want:  '7',
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got := testCase.in.highCard(testCase.joker)
			is.Equal(string(got), string(testCase.want))
		})
	}
}

func TestSecondOrderingRule(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in1, in2 Hand
		want     bool
	}{
		"first": {
			in1:  Hand{cards: [5]Card{'3', '3', '3', '3', '2'}},
			in2:  Hand{cards: [5]Card{'2', 'A', 'A', 'A', 'A'}},
			want: true,
		},
		"second": {
			in1:  Hand{cards: [5]Card{'7', '7', '8', '8', '8'}},
			in2:  Hand{cards: [5]Card{'7', '7', '7', '8', '8'}},
			want: true,
		},
		"flip first": {
			in1:  Hand{cards: [5]Card{'2', 'A', 'A', 'A', 'A'}},
			in2:  Hand{cards: [5]Card{'3', '3', '3', '3', '2'}},
			want: false,
		},
		"flip second": {
			in1:  Hand{cards: [5]Card{'7', '7', '7', '8', '8'}},
			in2:  Hand{cards: [5]Card{'7', '7', '8', '8', '8'}},
			want: false,
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got := testCase.in1.secondOrderingRule(testCase.in2, false)
			is.Equal(got, testCase.want)
		})
	}
}

func TestHandStrongerType(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in1, in2 string
		want     Hand
	}{
		"first": {
			in1:  "33332",
			in2:  "2AAAA",
			want: Hand{[5]Card{'3', '3', '3', '3', '2'}, FourOfAKind},
		},
		"second": {
			in1:  "77888",
			in2:  "77788",
			want: Hand{[5]Card{'7', '7', '8', '8', '8'}, FullHouse},
		},
		"flip first": {
			in1:  "2AAAA",
			in2:  "33332",
			want: Hand{[5]Card{'3', '3', '3', '3', '2'}, FourOfAKind},
		},
		"flip second": {
			in1:  "77788",
			in2:  "77888",
			want: Hand{[5]Card{'7', '7', '8', '8', '8'}, FullHouse},
		},

		"troubleshoot": {
			in1:  "KTJJT",
			in2:  "T55J5",
			want: Hand{[5]Card{'T', '5', '5', 'J', '5'}, ThreeOfAKind},
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			h1, err := parseHand(testCase.in1)
			is.NoErr(err)

			h2, err := parseHand(testCase.in2)
			is.NoErr(err)

			got := h1.StrongerThan(h2, false)
			is.Equal(got, testCase.want)
		})
	}
}

func TestParseHand(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want Hand
	}{
		"first": {
			in:   "33332",
			want: Hand{[5]Card{'3', '3', '3', '3', '2'}, FourOfAKind},
		},
		"second": {
			in:   "77888",
			want: Hand{[5]Card{'7', '7', '8', '8', '8'}, FullHouse},
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := parseHand(testCase.in)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}

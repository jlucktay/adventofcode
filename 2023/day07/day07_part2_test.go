/*
Now, the above example goes very differently:

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483

    32T3K is still the only one pair; it doesn't contain any jokers, so its strength doesn't increase.
    KK677 is now the only two pair, making it the second-weakest hand.
    T55J5, KTJJT, and QQQJA are now all four of a kind! T55J5 gets rank 3, QQQJA gets rank 4, and KTJJT gets rank 5.

With the new joker rule, the total winnings in this example are 5905.
*/

package main

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestPart2(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   []string
		want int
	}{
		"empty": {[]string{}, 0},
		"input": {
			in:   strings.Split(camelCards, "\n"),
			want: 5905,
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := Part2(testCase.in)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}

func TestApplyJokerRule(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want string
	}{
		"input 1": {
			in:   "32T3K",
			want: "32T3K (OnePair)",
		},
		"input 1 but jokers": {
			in:   "J2TJK",
			want: "K2TKK (ThreeOfAKind)",
		},
		"input 2": {
			in:   "KK677",
			want: "KK677 (TwoPair)",
		},
		"input 3": {
			in:   "T55J5",
			want: "T5555 (FourOfAKind)",
		},
		"input 4": {
			in:   "KTJJT",
			want: "KTTTT (FourOfAKind)",
		},
		"input 5": {
			in:   "QQQJA",
			want: "QQQQA (FourOfAKind)",
		},

		"all jokers": {
			in:   "JJJJJ",
			want: "JJJJJ (FiveOfAKind)",
		},
		"4 jokers": {
			in:   "JJJJA",
			want: "AAAAA (FiveOfAKind)",
		},
		"3 jokers and other two are same": {
			in:   "JJJAA",
			want: "AAAAA (FiveOfAKind)",
		},
		"3 jokers and other two differ": {
			in:   "JJJA2",
			want: "AAAA2 (FourOfAKind)",
		},
		"2 jokers and other three are all same": {
			in:   "JJAAA",
			want: "AAAAA (FiveOfAKind)",
		},
		"2 jokers and another pair and last card differs": {
			in:   "JJAA2",
			want: "AAAA2 (FourOfAKind)",
		},
		"2 jokers and other three all differ": {
			in:   "JJA23",
			want: "AAA23 (ThreeOfAKind)",
		},
		"1 joker and other four are all same": {
			in:   "JAAAA",
			want: "AAAAA (FiveOfAKind)",
		},
		"1 joker and two pair": {
			in:   "JA2A2",
			want: "AA2A2 (FullHouse)",
		},
		"1 joker and one pair and other two differ": {
			in:   "JAA34",
			want: "AAA34 (ThreeOfAKind)",
		},
		"1 joker and other four all differ": {
			in:   "JA234",
			want: "AA234 (OnePair)",
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := parseHand(testCase.in)
			is.NoErr(err)

			is.Equal(got.applyJokerRule().String(), testCase.want)
		})
	}
}

func TestStrongerThanJoker(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in, against string
		want        bool
	}{
		// 32T3K is still the only one pair; it doesn't contain any jokers, so its strength doesn't increase.
		"32T3K": {
			in:      "32T3K",
			against: "T55J5",
			want:    false,
		},

		// KK677 is now the only two pair, making it the second-weakest hand.
		"KK677 against lower rank": {
			in:      "KK677",
			against: "32T3K",
			want:    true,
		},
		"KK677 against higher rank 2": {
			in:      "KK677",
			against: "T55J5",
			want:    false,
		},

		// T55J5, KTJJT, and QQQJA are now all four of a kind!
		// T55J5 gets rank 3, QQQJA gets rank 4, and KTJJT gets rank 5.
		"T55J5 against lower rank 1": {
			in:      "T55J5",
			against: "32T3K",
			want:    true,
		},
		"T55J5 against lower rank 2": {
			in:      "T55J5",
			against: "KK677",
			want:    true,
		},
		"T55J5 against higher rank 4": {
			in:      "T55J5",
			against: "QQQJA",
			want:    false,
		},

		// T55J5 gets rank 3, QQQJA gets rank 4, and KTJJT gets rank 5.
		"QQQJA against lower rank 1": {
			in:      "QQQJA",
			against: "32T3K",
			want:    true,
		},
		"QQQJA against lower rank 2": {
			in:      "QQQJA",
			against: "KK677",
			want:    true,
		},
		"QQQJA against lower rank 3": {
			in:      "QQQJA",
			against: "T55J5",
			want:    true,
		},
		"QQQJA against higher rank 5": {
			in:      "QQQJA",
			against: "KTJJT",
			want:    false,
		},

		// T55J5 gets rank 3, QQQJA gets rank 4, and KTJJT gets rank 5.
		"KTJJT against lower rank 1": {
			in:      "KTJJT",
			against: "32T3K",
			want:    true,
		},
		"KTJJT against lower rank 2": {
			in:      "KTJJT",
			against: "KK677",
			want:    true,
		},
		"KTJJT against lower rank 3": {
			in:      "KTJJT",
			against: "T55J5",
			want:    true,
		},
		"KTJJT against lower rank 4": {
			in:      "KTJJT",
			against: "QQQJA",
			want:    true,
		},

		// rank 991 for hand 7JJ77 (FullHouse) (joker version: 77777 (FiveOfAKind) )
		// rank 992 for hand 7777J (FourOfAKind) (joker version: 77777 (FiveOfAKind) )
		"this seems off": {
			in:      "7JJ77",
			against: "7777J",
			want:    false,
		},

		"literally from the spec": {
			in:      "JKKK2",
			against: "QQQQ2",
			want:    false,
		},
		"literally from the spec - inverted": {
			in:      "QQQQ2",
			against: "JKKK2",
			want:    true,
		},

		// falling through to second order rule for 7JJ77 (FullHouse) vs J6666 (FourOfAKind) ; jokers: true
		// returning 7JJ77 (FullHouse)
		//
		// second order should come into play here
		// return the full house because it starts with a 7 and not a J
		"from part 2 input": {
			in:      "7JJ77",
			against: "J6666",
			want:    true,
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase

		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			in, err := parseHand(testCase.in)
			is.NoErr(err)

			against, err := parseHand(testCase.against)
			is.NoErr(err)

			got := in.StrongerThan(against, true)
			if testCase.want {
				is.Equal(got, in)
			} else {
				is.Equal(got, against)
			}
		})
	}
}

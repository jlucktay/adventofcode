/*
The escalator doesn't move. The Elf explains that it probably needs more joltage to overcome the static friction of the system and hits the big red "joltage limit safety override" button. You lose count of the number of times she needs to confirm "yes, I'm sure" and decorate the lobby a bit while you wait.

Now, you need to make the largest joltage by turning on exactly twelve batteries within each bank.

The joltage output for the bank is still the number formed by the digits of the batteries you've turned on; the only difference is that now there will be 12 digits in each bank's joltage output instead of two.

Consider again the example from before:

987654321111111
811111111111119
234234234234278
818181911112111

Now, the joltages are much larger:

    In 987654321111111, the largest joltage can be found by turning on everything except some 1s at the end to produce 987654321111.
    In the digit sequence 811111111111119, the largest joltage can be found by turning on everything except some 1s, producing 811111111119.
    In 234234234234278, the largest joltage can be found by turning on everything except a 2 battery, a 3 battery, and another 2 battery near the start to produce 434234234278.
    In 818181911112111, the joltage 888911112111 is produced by turning on everything except some 1s near the front.

The total output joltage is now much larger: 987654321111 + 811111111119 + 434234234278 + 888911112111 = 3121910778619.

What is the new total output joltage?
*/

package day03_test

import (
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/2025/day03"
)

func TestPart2(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in   string
		want int64
	}{
		"empty":  {"", 0},
		"line 1": {`987654321111111`, 987_654_321_111},
		"line 2": {`811111111111119`, 811_111_111_119},
		"line 3": {`234234234234278`, 434_234_234_278},
		"line 4": {`818181911112111`, 888_911_112_111},
		"for example": {`987654321111111
811111111111119
234234234234278
818181911112111
`, 3_121_910_778_619},
	}

	for desc, testCase := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := day03.Part2(testCase.in)
			is.NoErr(err)
			is.Equal(got, testCase.want)
		})
	}
}

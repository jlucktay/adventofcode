package day13

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/matryer/is"
)

const INPUT = `[[[10],7,3,[[6,4,4,1],[],[2,6,7,10]],6],[]]`

const INPUT_EXPLODED = `
[
	[
		[
			10
		],
		7,
		3,
		[
			[
				6,
				4,
				4,
				1
			],
			[],
			[
				2,
				6,
				7,
				10
			]
		],
		6
	],
	[]
]`

func TestParseList(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		input  string
		output *PacketList
	}{
		"smol": {
			input: `[42]`,
			output: &PacketList{
				PacketInteger(42),
			},
		},
		"given input": {
			input: INPUT,
			output: &PacketList{
				&PacketList{
					&PacketList{
						PacketInteger(10),
					},
					PacketInteger(7),
					PacketInteger(3),
					&PacketList{
						&PacketList{
							PacketInteger(6),
							PacketInteger(4),
							PacketInteger(4),
							PacketInteger(1),
						},
						&PacketList{},
						&PacketList{
							PacketInteger(2),
							PacketInteger(6),
							PacketInteger(7),
							PacketInteger(10),
						},
					},
					PacketInteger(6),
				},
				&PacketList{},
			},
		},
	}

	for desc, testCase := range testCases {
		desc, testCase := desc, testCase
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			actual, _, err := parseList(testCase.input, 0)
			is.NoErr(err)

			t.Logf("expected: %#v", testCase.output)
			t.Logf("actual:   %#v", actual)

			if diff := cmp.Diff(testCase.output, actual); diff != "" {
				t.Errorf("parseList() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

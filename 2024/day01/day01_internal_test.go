package day01

import (
	"errors"
	"strconv"
	"testing"

	"github.com/matryer/is"
)

func TestParseInput(t *testing.T) {
	is := is.New(t)

	testCases := map[string]struct {
		in      string
		want    twoLists
		wantErr error
	}{
		"empty": {"", twoLists{left: []int{}, right: []int{}}, nil},
		"for example": {
			`3   4
4   3
2   5
1   3
3   9
3   3
`,
			twoLists{
				left:  []int{3, 4, 2, 1, 3, 3},
				right: []int{4, 3, 5, 3, 9, 3},
			},
			nil,
		},
		"error left": {
			"a 1",
			twoLists{},
			&strconv.NumError{
				Func: "Atoi",
				Num:  "a",
				Err:  errors.New("invalid syntax"),
			},
		},
		"error right": {
			"1 a",
			twoLists{},
			&strconv.NumError{
				Func: "Atoi",
				Num:  "a",
				Err:  errors.New("invalid syntax"),
			},
		},
	}

	for desc, testCase := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got, err := parseInput(testCase.in)
			is.Equal(err, testCase.wantErr)
			is.Equal(got, testCase.want)
		})
	}
}

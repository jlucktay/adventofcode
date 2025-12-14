package slice_test

import (
	"slices"
	"testing"

	"github.com/matryer/is"

	"go.jlucktay.dev/adventofcode/crunchy/slice"
)

func TestUnionAny(t *testing.T) {
	is := is.New(t)

	type testCase[T any] struct {
		a, b, want []T
	}

	testCases := map[string]testCase[any]{
		"nil":   {a: nil, b: nil, want: nil},
		"a nil": {a: nil, b: []any{'a', 2, "D"}, want: []any{'a', 2, "D"}},
		"b nil": {a: []any{'a', 2, "D"}, b: nil, want: []any{'a', 2, "D"}},
		"mix":   {a: []any{'a'}, b: []any{2}, want: []any{'a', 2}},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got := slice.Union(tc.a, tc.b)

			if tc.a != nil && tc.b != nil {
				for _, aa := range tc.a {
					is.True(slices.Contains(got, aa))
				}
				for _, bb := range tc.b {
					is.True(slices.Contains(got, bb))
				}
			} else {
				is.Equal(tc.want, got)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	is := is.New(t)

	type testCase[T any] struct {
		input, want []T
		indices     []int
	}

	testCases := map[string]testCase[string]{
		"nil":               {input: nil, want: nil, indices: nil},
		"nil input":         {input: nil, want: nil, indices: []int{42}},
		"nil indices":       {input: []string{"42"}, want: []string{"42"}, indices: nil},
		"first index":       {input: []string{"6", "7", "27", "42"}, want: []string{"7", "27", "42"}, indices: []int{0}},
		"last index":        {input: []string{"6", "7", "27", "42"}, want: []string{"6", "7", "27"}, indices: []int{3}},
		"out of bounds":     {input: []string{"6", "7", "27", "42"}, want: []string{"6", "7", "27", "42"}, indices: []int{27, 42}},
		"mix in and out":    {input: []string{"6", "7", "27", "42"}, want: []string{"6", "7", "42"}, indices: []int{4, 2, 4, 2}},
		"delete everything": {input: []string{"6", "7", "27", "42"}, want: []string{}, indices: []int{0, 1, 2, 3}},
		"negative index":    {input: []string{"6", "7", "27", "42"}, want: []string{"6", "7", "27", "42"}, indices: []int{-5}},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			is := is.New(t)

			got := slice.Delete(tc.input, tc.indices...)

			is.True(slices.Equal(got, tc.want))
		})
	}
}

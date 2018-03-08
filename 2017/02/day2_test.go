package day2

import (
	"reflect"
	"testing"
)

func TestConvertInput(t *testing.T) {
	cases := []struct {
		in   string
		want [][]int
	}{
		{
			`1
2	3
4	5	a`,
			[][]int{
				{1},
				{2, 3},
				{4, 5, 0},
			},
		},
		{
			`1
2	3
4	5	6`,
			[][]int{
				{1},
				{2, 3},
				{4, 5, 6},
			},
		},
		{
			`5	1	9	5
7	5	3
2	4	6	8`,
			[][]int{
				{5, 1, 9, 5},
				{7, 5, 3},
				{2, 4, 6, 8},
			},
		},
		{
			`5	9	2	8
9	4	7	3
3	8	6	5`,
			[][]int{
				{5, 9, 2, 8},
				{9, 4, 7, 3},
				{3, 8, 6, 5},
			},
		},
	}

	for _, c := range cases {
		got := ConvertInput(c.in)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ConvertInput(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

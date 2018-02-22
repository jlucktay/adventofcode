/*
When fetching input from text files, we need it converted into a more useful type, i.e. a 2-d slice.
*/

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
			`5	9	2	8
9	4	7	3
3	8	6	5`,
			[][]int{
				{5, 9, 2, 8},
				{9, 4, 7, 3},
				{3, 8, 6, 5},
				{},
			},
		},
		{
			"1",
			[][]int{
				{1, 0, 0, 0},
				{},
				{},
				{},
			},
		},
		{
			`0	1	2	3
4	5	6	7`,
			[][]int{
				{0, 1, 2, 3},
				{4, 5, 6, 7},
				{},
				{},
			},
		},
	}

	for _, c := range cases {
		got := ConvertInput(c.in)

		for i := range got {
			if !reflect.DeepEqual(got[i], c.want[i]) {
				t.Errorf("ConvertInput(%v) == %v, want %v", c.in, got[i], c.want[i])
			}
		}
	}
}

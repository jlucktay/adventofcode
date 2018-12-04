package aoc201801

import (
	"reflect"
	"testing"
)

func TestProcessInput(t *testing.T) {
	cases := []struct {
		in   string
		want []int
	}{
		{
			`+1
-1
+2
-2
+3
-3`,
			[]int{1, -1, 2, -2, 3, -3},
		},
		{
			`+1
+-1
+2`,
			[]int{1, 2},
		},
	}

	for _, c := range cases {
		got := processInput(c.in)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("f(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

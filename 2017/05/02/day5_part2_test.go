/*
For example, consider the following list of jump offsets:

0
3
0
1
-3

Positive jumps ("forward") move downward; negative jumps move upward.

...

Now, the jumps are even stranger: after each jump, if the offset was three or more,
instead decrease it by 1. Otherwise, increase it by 1 as before.

Using this rule with the above example, the process now takes 10 steps, and the
offset values after finding the exit are left as 2 3 2 3 -1.
*/

package main

import (
	"reflect"
	"testing"
)

func TestFunction(t *testing.T) {
	cases := []struct {
		in        []int
		wantSteps int
		wantSlice []int
	}{
		{
			[]int{0, 3, 0, 1, -3},
			10,
			[]int{2, 3, 2, 3, -1}},
	}

	for _, c := range cases {
		originalInput := make([]int, len(c.in))
		copy(originalInput, c.in)
		got := followJumpOffsets(c.in)

		if got != c.wantSteps {
			t.Errorf("f(%v) == %v, want %v", originalInput, got, c.wantSteps)
		}

		if !reflect.DeepEqual(c.in, c.wantSlice) {
			t.Errorf("Input slice '%v' does not match desired end state: %v", c.in, c.wantSlice)
		}
	}
}

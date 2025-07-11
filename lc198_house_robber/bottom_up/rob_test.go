package lc198bottomup

import (
	"testing"
	"fmt"
)

func TestRob(t *testing.T) {
	testCases := []struct {
		input  []int
		expect int
	}{
		{input: []int{1, 2, 3, 1}, expect: 4},
		{input: []int{2, 7, 9, 3, 1}, expect: 12},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			res := rob(tc.input)
			if res != tc.expect {
				t.Errorf("For input %v expected %v, got %v", tc.input, tc.expect, res)
			}
		})
	}
}

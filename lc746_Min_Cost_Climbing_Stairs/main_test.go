package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	testCases := []struct {
		args []int
		want int
	}{
		{
			args: []int{10, 15, 20},
			want: 15,
		},
		{
			args: []int{15, 10, 20},
			want: 10,
		},
		{
			args: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := minCostClimbingStairs(tc.args)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Expected %v, got %v, args %v", tc.want, got, tc.args)
			}
		})
	}
}

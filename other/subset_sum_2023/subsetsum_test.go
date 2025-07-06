package main

import "testing"

func TestIsSubsetSum(t *testing.T) {
	tests := []struct {
		set  []int
		sum  int
		want bool
	}{
		{
			set:  []int{3, 34, 4, 12, 5, 2},
			sum:  9,
			want: true,
		},
		{
			set:  []int{3, 34, 4, 12, 5, 2},
			sum:  30,
			want: false,
		},
		{
			set:  []int{3, 34, 4, 12, 5, 2},
			sum:  100,
			want: false,
		},
		{
			set:  []int{},
			sum:  0,
			want: true,
		},
		{
			set:  []int{25},
			sum:  0,
			want: true,
		},
		{
			set:  []int{25},
			sum:  25,
			want: true,
		},
		{
			set:  []int{5, 25},
			sum:  25,
			want: true,
		},
		{
			set:  []int{25},
			sum:  30,
			want: false,
		},
		{
			set:  []int{1, 9, 17, 3, 5, 3, 2},
			sum:  16,
			want: true,
		},
		{
			set:  []int{1, 9, 17, 3, 5, 3, 2},
			sum:  19,
			want: true,
		},
		{
			set:  []int{1, 9, 17, 3, 5, 3, 2},
			sum:  20,
			want: true,
		},
		{
			set:  []int{1, 9, 17, 3, 5, 3, 2},
			sum:  4,
			want: true,
		},
		{
			set:  []int{7, 6, 10},
			sum:  12,
			want: false,
		},
	}

	for _, tt := range tests {
		got := isSubsetSum(tt.set, tt.sum)
		if got != tt.want {
			t.Errorf("isSubsetSum(%v, %d) = %v; want %v", tt.set, tt.sum, got, tt.want)
		}
	}
}

package main

import (
	"fmt"
	"testing"
)

func Test_hammingDistance(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{x: 1, y: 4}, want: 2},
		{args: args{x: 0, y: 0}, want: 0},       // Both numbers are same, distance is 0
		{args: args{x: 15, y: 0}, want: 4},      // 15 is 1111 in binary, distance is 4
		{args: args{x: 2, y: 3}, want: 1},       // 2 is 10 and 3 is 11 in binary, distance is 1
		{args: args{x: 3, y: 5}, want: 2},       // 3 is 011 and 5 is 101 in binary, distance is 2
		{args: args{x: 8, y: 16}, want: 2},      // Different bits
		{args: args{x: 31, y: 16}, want: 4},     // 31 is 11111 and 16 is 10000, distance is 4
		{args: args{x: 1024, y: 2048}, want: 2}, // Differ by 2 bits
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := hammingDistance(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("hammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

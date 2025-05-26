package lc2000

import (
	"reflect"
	"testing"
)

func Test_minimumAbsDifference(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "[4,2,1,3]", args: args{arr: []int{4, 2, 1, 3}}, want: [][]int{{1, 2}, {2, 3}, {3, 4}}},
		{name: "[1,3,6,10,15]", args: args{arr: []int{1, 3, 6, 10, 15}}, want: [][]int{{1, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAbsDifference(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumAbsDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

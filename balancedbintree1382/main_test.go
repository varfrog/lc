package main

import (
	"reflect"
	"testing"
)

func Test_balanceBST_NoBalancingNeeded(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Just root",
			args: args{
				root: &TreeNode{},
			},
			want: &TreeNode{},
		},
		{
			name: "Already balanced",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balanceBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("balanceBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_balanceBST_BalanceCase1(t *testing.T) {
	got := *balanceBST(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val:   4,
					Right: nil,
				},
			},
		},
	})

	expectation1 := reflect.DeepEqual(got, TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	})

	expectation2 := reflect.DeepEqual(got, TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	})

	if !expectation1 && !expectation2 {
		t.Errorf("balanceBST() = %v, not matching expected neither two", got)
	}
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortArray(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		//{name: "empty", input: []int{}, expected: []int{}},
		//{name: "single element", input: []int{1}, expected: []int{1}},
		//{name: "already sorted", input: []int{1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "reverse order", input: []int{3, 2, 1}, expected: []int{1, 2, 3}},
		{name: "negatives and positives", input: []int{5, -2, 0, 3, 9}, expected: []int{-2, 0, 3, 5, 9}},
		{name: "all duplicates", input: []int{3, 3, 3}, expected: []int{3, 3, 3}},
		{name: "shuffled 0-9", input: []int{9, 1, 8, 2, 7, 3, 6, 4, 5, 0}, expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{
			name:     "mixed with duplicates",
			input:    []int{3, 2, 1, -100, 0, 0, 4, 5, 6, 1, 7, 8, 9, 3},
			expected: []int{-100, 0, 0, 1, 1, 2, 3, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			result := sortArray(append([]int(nil), tc.input...)) // avoid mutating original
			assert.Equal(t, tc.expected, result)
		})
	}
}

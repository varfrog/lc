package v1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_sortArray(t *testing.T) {
	arrays := [][]int{
		{1},
		{1, 2, 3},
		{1, 3, 2},
		{3, 2, 1},
		{3, 2, 1, 1, 7},
		{3, 2, 1, -100, 0, 0, 4, 5, 6, 1, 7, 8, 9, 3},
	}
	for _, arr := range arrays {
		assert.Truef(t, sort.IntsAreSorted(sortArray(arr)), fmt.Sprintf("%v not sorted", arr))
	}
}

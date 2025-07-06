package main

import (
	"fmt"
)

func isSubsetSum(arr []int, t int) bool {
	if len(arr) == 0 {
		if t == 0 {
			return true
		}
		return false
	}
	if t == 0 {
		return true
	}

	// dp is a memoization matrix where dp[i][j]=true iff there is a subset in arr[0..j) (of elements from index 0 to i
	// inclusively) that sums up to j.
	dp := make([][]bool, len(arr)+1) // +1 for empty subset

	// Init j cols
	numCols := t + 1 // +1 for zero
	for i := range dp {
		dp[i] = make([]bool, numCols)
	}

	// Leftmost column. Every set has an empty subset which sums to 0 (j=0).
	for i := 0; i < len(dp); i++ {
		dp[i][0] = true
	}

	// First row. No empty set sums to a non-empty j.
	for j := 1; j < numCols; j++ {
		dp[0][j] = false
	}

	// Main algorithm
	for i := 1; i < len(dp); i++ {
		for j := 1; j < numCols; j++ {
			curVal := arr[i-1]
			if curVal > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || // Exclude current element
					dp[i-1][j-curVal] // Is there a subset in arr at indices [0; i-1] that sums to j-curVal?
			}
		}
	}

	return dp[len(dp)-1][t]
}

func main() {
	arr := []int{5}
	sum := 5

	if isSubsetSum(arr, sum) {
		fmt.Println("Found a subset with the given sum")
	} else {
		fmt.Println("No subset with the given sum found")
	}
}

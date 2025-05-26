package lc2000

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	// Find the min distance
	minDistance := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		dist := arr[i] - arr[i-1]
		if dist < minDistance {
			minDistance = dist
		}
	}

	// Find pairs with distance = minDistance
	var resultPairs [][]int
	for i := 1; i < len(arr); i++ {
		dist := arr[i] - arr[i-1]
		if dist == minDistance {
			resultPairs = append(resultPairs, []int{arr[i-1], arr[i]})
		}
	}

	return resultPairs
}

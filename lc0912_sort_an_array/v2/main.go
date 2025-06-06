package main

import "fmt"

func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	nums = quicksort(nums)
	return nums
}

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left := make([]int, 0, len(nums))
	right := make([]int, 0, len(nums))
	pivotGroup := make([]int, 0) // To store the pivot. May include duplicates if it appears multiple times in nums
	pivot := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			left = append(left, nums[i])
		} else if nums[i] > pivot {
			right = append(right, nums[i])
		} else {
			pivotGroup = append(pivotGroup, nums[i])
		}
	}

	leftSorted := quicksort(left)
	rightSorted := quicksort(right)

	result := append(leftSorted, pivotGroup...)
	result = append(result, rightSorted...)

	return result
}

func main() {
	nums := []int{10, 5, 13, 2}
	sortArray(nums)
	fmt.Printf("%d\n", nums)
}

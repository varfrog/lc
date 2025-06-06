package v1

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	return *quicksort(&nums, 0, len(nums)-1)
}

func quicksort(nums *[]int, low, high int) *[]int {
	if high < low {
		return nil
	}

	partitionIndex := partition(nums, low, high)
	quicksort(nums, low, partitionIndex-1)
	quicksort(nums, partitionIndex+1, high)

	return nums
}

func partition(nums *[]int, lowIndex, highIndex int) (i int) {
	pivot := (*nums)[highIndex]
	i = lowIndex
	for j := lowIndex; j < highIndex; j++ {
		if (*nums)[j] < pivot {
			(*nums)[j], (*nums)[i] = (*nums)[i], (*nums)[j]
			i++
		}
	}
	(*nums)[highIndex], (*nums)[i] = (*nums)[i], (*nums)[highIndex]

	return i
}

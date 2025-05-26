package kata

import (
	"fmt"
	"sort"
	"sync"
)

/*
John and Mary want to travel between a few towns A, B, C ... Mary has on a sheet of paper a list of distances between
these towns. ls = [50, 55, 57, 58, 60]. John is tired of driving, and he says to Mary that he doesn't want to drive more
than t = 174 miles, and he will visit only k = 3 towns.

Which distances, hence which towns, they will choose so that the sum of the distances is the biggest possible to please
Mary and John?
*/

func ChooseBestSum(t, k int, ls []int) int {
	lsLen := len(ls)
	if lsLen < k {
		return -1
	}

	var sums []int
	for _, mask := range buildArrayMasks(lsLen, k) {
		sums = append(sums, sumLs(applyArrayMask(ls, mask)))
	}
	sort.Ints(sums)

	result := sums[0]
	for i := 1; i < len(sums); i++ {
		if sums[i] > t {
			break
		}
		result = sums[i]
	}

	if result > t {
		return -1
	}
	return result
}

func sumLs(ls []int) (result int) {
	for _, val := range ls {
		result += val
	}
	return
}

// buildArrayMasks returns all possible combinations of []int8 with elements 0 and 1, having k number of 1's.
func buildArrayMasks(lsLen int, k int) (masks [][]int8) {
	resCh := make(chan []int8)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for m := range resCh {
			if countOnes(m) == k {
				masks = append(masks, m)
			}
		}
		wg.Done()
	}()
	buildArrayMasksAsync(make([]int8, lsLen), resCh)
	wg.Wait()

	return masks
}

// forkArrayMask is a wrapper for forkArrayMask
func buildArrayMasksAsync(mask []int8, res chan<- []int8) {
	res <- mask
	wg := sync.WaitGroup{}
	forkArrayMask(mask, 0, res, &wg)
	wg.Wait()
	close(res)
}

func forkArrayMask(mask []int8, forkAtIndex int, res chan<- []int8, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	maskLen := len(mask)
	newMask := make([]int8, maskLen)
	copy(newMask, mask)

	forkValue := int8(0)
	if newMask[forkAtIndex] == forkValue {
		forkValue = 1
	}
	newMask[forkAtIndex] = forkValue
	res <- newMask

	if forkAtIndex+1 == maskLen {
		return
	}

	forkArrayMask(mask, forkAtIndex+1, res, wg)
	forkArrayMask(newMask, forkAtIndex+1, res, wg)
}

// countOnes count 1's in a mask, e.g. given [0 1 0 1 1], the result would be 3.
func countOnes(mask []int8) (result int) {
	for _, val := range mask {
		if val == 1 {
			result++
		}
	}
	return
}

// applyArrayMask works like a bit mask. It returns a copy of ls with each nth element zeroed out if the nth element in
// mask is zero. E.g:
// ls =     [1 2 3 4 5]
// mask =   [0 1 0 0 1]
// result = [0 2 0 0 5]
func applyArrayMask(ls []int, mask []int8) []int {
	lsLen := len(ls)
	result := make([]int, lsLen)
	for i := 0; i < lsLen; i++ {
		if mask[i] != 0 {
			result[i] = ls[i]
		} else {
			result[i] = 0
		}
	}

	return result
}

func Run() {
	fmt.Println(ChooseBestSum(173, 3, []int{50, 55, 57, 58, 60}))
}

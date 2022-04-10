package kata

/*
John and Mary want to travel between a few towns A, B, C ... Mary has on a sheet of paper a list of distances between
these towns. ls = [50, 55, 57, 58, 60]. John is tired of driving, and he says to Mary that he doesn't want to drive more
than t = 174 miles, and he will visit only k = 3 towns.

Which distances, hence which towns, they will choose so that the sum of the distances is the biggest possible to please
Mary and John?
*/

func ChooseBestSum(t, k int, ls []int) int {
	if len(ls) < k {
		return -1
	}
	return 0
}

// applyArrayMask works like a bit mask. It returns a copy of ls with each nth element zeroes out if the nth element in
// mask is zero. E.g:
// ls = 	[1 2 3 4 5]
// mask = 	[0 1 0 0 1]
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

package lc198topdown

func rob(nums []int) int {
	// memo holds the value of recurse(), key i is a slice nums of length i.
	memo := make(map[int]int, len(nums)+1)

	return recurse(nums, 0, memo)
}

func recurse(nums []int, start int, memo map[int]int) int {
	if start > len(nums)-1 {
		return 0
	}
	if val, ok := memo[start]; ok {
		return val
	}

	robThis := nums[start] + recurse(nums, start+2, memo)
	robNext := recurse(nums, start+1, memo)

	memo[start] = max(robThis, robNext)

	return memo[start]
}

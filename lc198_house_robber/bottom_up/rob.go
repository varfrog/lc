package lc198bottomup

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	// dp[a] = max money arriving to rob at house index a. Can skip house index a.
	dp := make([]int, n)

	// Base cases
	dp[n-1] = nums[n-1]
	dp[n-2] = max(nums[n-2], nums[n-1])
	dp[n-3] = max(nums[n-3]+nums[n-1], nums[n-2])

	// Build DP
	for i := n - 4; i >= 0; i-- {
		dp[i] = max(nums[i]+dp[i+2], dp[i+1])
	}

	// Get the max
	runningMax := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] > runningMax {
			runningMax = dp[i]
		}
	}

	return runningMax
}

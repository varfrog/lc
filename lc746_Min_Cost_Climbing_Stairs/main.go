package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost) // n = number of staircases.
	if n == 2 {
		return min(cost[0], cost[1])
	}

	// dp[i] - minimum cost to climb to the top, starting from the ith staircase.
	dp := make([]int, n)

	// Steps are numbered 0 to n-1. Final pos is n (the top).
	//
	// When we're at the top (n), we got here either:
	// A: paying the cost for staircase n-1 and taking one step,
	// B: paying the cost for staircase n-2 and taking two steps.
	// This is the base for dp.
	dp[n-1] = cost[n-1] // Cost to reach the top from step n-1
	dp[n-2] = cost[n-2] // Cost to reach the top from step n-2

	// When we're at the third to last step (n-3), our minimum cost to reach the top is:
	// 1. Pay the cost for staircase n-3.
	// 2. Either:
	// 	  A: move to staircase n-2 and pay dp[n-2]
	//    B: move to staircase n-1 and pay dp[n-1]
	// So dp[n-3] = cost[n-3] + min(dp[n-1], dp[n-2]).
	// And so it repeats until staircase 0.

	for i := n - 3; i >= 0; i-- {
		dp[i] = cost[i] + min(dp[i+2], dp[i+1])
	}

	return min(dp[0], dp[1])
}

func main() {

}

package main

import "fmt"

/*
varfrog:
_ _ X <- here's our n
_ X X <- (Comment A)
X X X <- (Comment B)

Comment A: to go one step further, there's only 1 possibility - take one step
Comment B: to go two steps further, you can take 1 step + 1 step, or 2 steps. Total 2 ways.
*/

func climbStairs(n int) int {
	if n < 2 {
		return n
	}

	// dp[i] = num steps taken to reach step i.
	dp := make([]int, n+1) // +1 since we index one-based

	// Base cases
	dp[1] = 1 // 1
	dp[2] = 2 // 1,1; 2

	// Build dp bottom-up
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + // 1 possibility - take one step
			// Even if we have 2 ways to climb from step 'i-2' to step 'i', we don't need to multiply dp[i-2]
			// by 2, since the possibilities for +1 step is included in dp[i-1]
			dp[i-2]
	}

	return dp[n]
}

func main() {
	fmt.Printf("climbStairs: %d\n", climbStairs(10))
}

func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	// (target + total) must be even and non-negative to have a solution
	if (target+total)%2 != 0 || target+total < 0 || abs(target) > total {
		return 0
	}

	subsetSum := (target + total) / 2
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, subsetSum+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j <= subsetSum; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] += dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][subsetSum]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
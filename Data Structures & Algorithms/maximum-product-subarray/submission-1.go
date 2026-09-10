func maxProduct(nums []int) int {
	dp := make([][2]int, len(nums))
	for i := range nums {
		dp[i] = [2]int{math.MaxInt, math.MinInt}
	}
	maxVal := nums[0]
	if nums[0] < 0 {
		dp[0][0] = nums[0]
	} else if nums[0] > 0 {
		dp[0][1] = nums[0]
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			dp[i][1] = nums[i]
			if dp[i-1][1] != math.MinInt {
				dp[i][1] = dp[i-1][1] * dp[i][1]
			}

			if dp[i-1][0] != math.MaxInt {
				dp[i][0] = dp[i-1][0] * nums[i]
			}
		} else if nums[i] < 0 {
			dp[i][0] = nums[i]
			if dp[i-1][1] != math.MinInt {
				dp[i][0] = dp[i-1][1] * dp[i][0]
			}

			if dp[i-1][0] != math.MaxInt {
				dp[i][1] = dp[i-1][0] * nums[i]
			}
		} else {
			maxVal = max(maxVal, 0)
		}
		maxVal = max(maxVal, dp[i][1])
	}

	return maxVal
}
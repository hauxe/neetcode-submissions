func canPartition(nums []int) bool {
	total := 0
	for i := range nums {
		total += nums[i]
	}
	if total%2 > 0 {
		return false
	}
	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for i := range nums {
		for j := target; j >= nums[i]; j-- {
			if dp[j-nums[i]] {
				dp[j] = true
			}
		}
	}
	return dp[target]
}
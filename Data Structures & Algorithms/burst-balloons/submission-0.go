func maxCoins(nums []int) int {
	n := len(nums)
	balloons := make([]int, n+2)
	balloons[0], balloons[n+1] = 1, 1
	copy(balloons[1:n+1], nums)

	dp := make([][]int, n+2)
	for i := range n + 2 {
		dp[i] = make([]int, n+2)
	}
	for gap := 2; gap < n+2; gap++ {
		for i := 0; i+gap < n+2; i++ {
			j := i + gap
			for k := i + 1; k < j; k++ {
				score := dp[i][k] + dp[k][j] + balloons[i]*balloons[k]*balloons[j]
				dp[i][j] = max(dp[i][j], score)
			}
		}
	}
	return dp[0][n+1]
}
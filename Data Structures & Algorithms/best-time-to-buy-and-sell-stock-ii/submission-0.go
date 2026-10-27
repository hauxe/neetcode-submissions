func maxProfit(prices []int) int {
	n := len(prices)
	hold, sold, rest := make([]int, n), make([]int, n), make([]int, n)
	hold[0] = -1 * prices[0]
	for i := 1; i < n; i++ {
		hold[i] = max(hold[i-1], rest[i-1]-prices[i], sold[i-1]-prices[i])
		sold[i] = hold[i-1] + prices[i]
		rest[i] = max(rest[i-1], sold[i-1])
	}
	return max(sold[n-1], rest[n-1])
}
func maxProfit(prices []int) int {
	hold := make([]int, len(prices))
	sold := make([]int, len(prices))
	rest := make([]int, len(prices))
	hold[0] = -1 * prices[0]
	sold[0] = 0
	rest[0] = 0
	for i := 1; i < len(prices); i++ {
		hold[i] = max(hold[i-1], rest[i-1]-prices[i])
		sold[i] = hold[i-1] + prices[i]
		rest[i] = max(rest[i-1], sold[i-1])
	}
	return max(sold[len(prices)-1], rest[len(prices)-1])
}
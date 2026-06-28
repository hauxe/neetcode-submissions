func maxProfit(prices []int) int {
	pivot := 0
	maxProfit := 0
	for i:=1;i<len(prices);i++ {
		if prices[i] < prices[pivot] {
			pivot = i
		} else {
			maxProfit = max(prices[i]-prices[pivot], maxProfit)
		}
	}
	return maxProfit
}

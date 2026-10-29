func subarraySum(nums []int, k int) int {
	m := map[int]int{0: 1} // the previous sum = 0 always exist for 1 time
	sum, result := 0, 0
	for i := range nums {
		sum += nums[i]
		result += m[sum-k]
		m[sum]++
	}
	return result
}
func getConcatenation(nums []int) []int {
	n := len(nums)
	result := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		if i < n {
			result[i] = nums[i]
		} else {
			result[i] = nums[i-n]
		}
	}
	return result
}
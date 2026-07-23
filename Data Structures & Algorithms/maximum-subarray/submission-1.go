func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a, b := nums[0], nums[1]
	maxVal := max(a, b, a+b)
	for i := 2; i < len(nums); i++ {
		if a < 0 {
			a, b = b, nums[i]
		} else {
			a, b = a+b, nums[i]
		}
		maxVal = max(maxVal, a, b, a+b)
	}
	return maxVal
}
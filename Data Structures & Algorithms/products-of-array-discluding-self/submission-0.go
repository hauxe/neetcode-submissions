func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	val := 1
	// prefix loop
	for i := 0; i < len(nums); i++ {
		result[i] = val
		val *= nums[i]
	}
	val = 1
	// suffix loop
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= val
		val *= nums[i]
	}
	return result
}
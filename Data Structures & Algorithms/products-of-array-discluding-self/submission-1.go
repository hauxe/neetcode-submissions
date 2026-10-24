func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	// left -> right loop
	val := 1
	for i := 0; i < n; i++ {
		result[i] = val
		val *= nums[i]
	}
	val = 1
	// right -> left loop
	for i := n - 1; i >= 0; i-- {
		result[i] *= val
		val *= nums[i]
	}
	return result
}
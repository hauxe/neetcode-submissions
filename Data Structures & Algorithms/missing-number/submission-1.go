func missingNumber(nums []int) int {
	missing := len(nums)
	for i, num := range nums {
		missing ^= (i ^ num)
	}
	return missing
}
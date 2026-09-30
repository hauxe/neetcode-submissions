func missingNumber(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	n := len(nums)
	return (n*(n+1)/2)-sum
}

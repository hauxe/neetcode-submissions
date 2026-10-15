func removeElement(nums []int, val int) int {
	pivot := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot--
		}
	}
	return pivot+1
}
func removeDuplicates(nums []int) int {
	pivot := 1
	for i := 1; i < len(nums); i++ {
		if nums[pivot-1] != nums[i] {
			nums[pivot], nums[i] = nums[i], nums[pivot]
			pivot++
			continue
		}
	}
	return pivot
}
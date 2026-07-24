func canJump(nums []int) bool {
	pos := len(nums) - 1
	for i := pos - 1; i >= 0; i-- {
		if nums[i] >= pos-i {
			pos = i
			continue
		}
	}
	return pos == 0
}
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
func reverse(nums []int) {
	for begin, end := 0, len(nums)-1; begin < end; begin, end, nums[begin], nums[end] = begin+1, end-1, nums[end], nums[begin] {
	}
}
func findMin(nums []int) int {
	begin, end := 0, len(nums)-1
	target := min(nums[begin], nums[end])
	for begin <= end {
		middle := (begin+end)/2
		if nums[middle] > target {
			begin = middle+1
		} else {
			target = nums[middle]
			end = middle-1
		}
	}
	return target
}

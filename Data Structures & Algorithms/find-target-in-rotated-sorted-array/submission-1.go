func search(nums []int, target int) int {
	begin, end := 0, len(nums)-1
	for begin <= end {
		middle := (begin + end) / 2
		if nums[middle] == target {
			return middle
		}
		if target < nums[middle] {
			if nums[middle] >= nums[begin] && target >= nums[begin] {
				end = middle - 1
			} else if nums[middle] <= nums[end] {
				end = middle - 1
			} else {
				begin = middle + 1
			}
		} else {
			if nums[middle] >= nums[begin] {
				begin = middle + 1
			} else if nums[middle] <= nums[end] && target <= nums[end] {
				begin = middle + 1
			} else {
				end = middle - 1
			}
		}
	}
	return -1
}
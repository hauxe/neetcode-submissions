func search(nums []int, target int) int {
	begin, end := 0, len(nums)-1
	for begin <= end {
		middle := (begin+end)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			begin = middle+1
		} else {
			end = middle-1
		}
	}
	return -1
}

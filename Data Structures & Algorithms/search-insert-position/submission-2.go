func searchInsert(nums []int, target int) int {
	begin, end := 0, len(nums)-1
	var middle int
	for begin <= end {
		middle = (begin + end) / 2
		switch {
		case nums[middle] < target:
			begin = middle + 1
		case nums[middle] > target:
			end = middle - 1
		default:
			return middle
		}
	}
	if nums[middle] < target {
		return middle + 1
	}
	return middle
}
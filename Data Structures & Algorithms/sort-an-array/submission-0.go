func sortArray(nums []int) []int {
	temp := make([]int, len(nums))
	mergeSort(nums, temp, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, temp []int, from, to int) {
	if to-from < 1 {
		return
	}
	mid := (from + to) / 2
	mergeSort(nums, temp, from, mid)
	mergeSort(nums, temp, mid+1, to)
	i, j, start := from, mid+1, from
	for ; i <= mid && j <= to; start++ {
		if nums[i] <= nums[j] {
			temp[start] = nums[i]
			i++
		} else {
			temp[start] = nums[j]
			j++
		}
	}
	for ; i <= mid; i, start = i+1, start+1 {
		temp[start] = nums[i]
	}
	for ; j <= to; j, start = j+1, start+1 {
		temp[start] = nums[j]
	}
	copy(nums[from:to+1], temp[from:to+1])
}
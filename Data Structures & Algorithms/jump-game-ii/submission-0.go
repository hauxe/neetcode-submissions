func jump(nums []int) int {
	result := make([]int, len(nums))
	pos := len(nums) - 1
	result[pos] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < pos-i {
			continue
		}
		result[i] = 1
		last := -1
		for j := pos + 1; j < len(nums) && j <= i+nums[i]; j++ {
			if result[j] == 1 {
				last = j
				result[pos] = 0
				result[j] = 0
			}

		}
		result[len(nums)-1] = 1
		if last >= 0 && last < len(nums)-1 {
			result[last] = 1
		}
		pos = i
	}
	count := 0
	for i := 0; i < len(result)-1; i++ {
		if result[i] == 1 {
			count++
		}
	}
	return count
}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	currentIdx := 0
	for currentIdx < len(nums)-2 {
		left, right := currentIdx+1, len(nums)-1
		for left < right {
			v := nums[currentIdx] + nums[left] + nums[right]
			if v == 0 {
				result = append(result, []int{nums[currentIdx], nums[left], nums[right]})
				for ; left < right && nums[left] == nums[left+1]; left++{}
				for ; left < right && nums[right] == nums[right-1]; right--{}
				left++
				right--
			} else if v < 0 {
				left++
			} else {
				right--
			}
		}
		for ; currentIdx < len(nums)-2 && nums[currentIdx] == nums[currentIdx+1]; currentIdx++{}
		currentIdx++
	}
	return result
}

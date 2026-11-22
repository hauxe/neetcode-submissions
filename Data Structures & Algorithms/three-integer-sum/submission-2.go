func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var results [][]int
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			switch {
			case sum < 0:
				left++
			case sum > 0:
				right--
			default:
				for ; left < right && nums[left+1] == nums[left]; left++ {
				}
				for ; left < right && nums[right-1] == nums[right]; right-- {
				}
				results = append(results, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}
		}
	}
	return results
}
func fourSum(nums []int, target int) [][]int {
	var results [][]int
	sort.Ints(nums)
	n := len(nums)
	for i := range n {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i + 1 && nums[j-1] == nums[j] {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				switch {
				case sum > target:
					right--
				case sum < target:
					left++
				default:
					results = append(results, []int{nums[i], nums[j], nums[left], nums[right]})
					for ; left < right && nums[left+1] == nums[left]; left++ {
					}
					for ; left < right && nums[right-1] == nums[right]; right-- {
					}
					left++
					right--
				}
			}
		}
	}
	return results
}
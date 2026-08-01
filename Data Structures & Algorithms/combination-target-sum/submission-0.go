func combinationSum(nums []int, target int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	var result [][]int
	for i := range nums {
		if nums[i] == target {
			result = append(result, []int{nums[i]})
		} else if nums[i] < target {
			sums := combinationSum(nums[i:], target-nums[i])
			if len(sums) > 0 {
				for j := range sums {
					sums[j] = append(sums[j], nums[i])
				}
				result = append(result, sums...)
			}
		}
	}
	return result
}
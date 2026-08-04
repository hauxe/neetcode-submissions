func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return doSubsetsWithDup(nums)
}

func doSubsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	prev := 0
	var result [][]int
	for i := 1; i <= len(nums); i++ {
		if i < len(nums) && nums[i] == nums[prev] {
			continue
		}
		fmt.Println(prev, i, nums[i:])
		index := i - prev
		subsets := doSubsetsWithDup(nums[i:])
		result = append(result, subsets...)
		for k := 1; k <= index; k++ {
			subsetsWithCurrent := make([][]int, len(subsets))
			for j := range subsets {
				subset := make([]int, len(subsets[j]))
				copy(subset, subsets[j])
				subsetsWithCurrent[j] = subset
			}
			for j := range subsetsWithCurrent {
				for l := 0; l < k; l++ {
					subsetsWithCurrent[j] = append(subsetsWithCurrent[j], nums[prev])
				}
			}
			result = append(result, subsetsWithCurrent...)
		}
		return result
	}
	return nil
}
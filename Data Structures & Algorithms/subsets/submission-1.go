func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	set := subsets(nums[1:])
	for i := range set {
		item := make([]int, len(set[i]))
		copy(item, set[i])
		item = append(item, nums[0])
		set = append(set, item)
	}
	return set
}
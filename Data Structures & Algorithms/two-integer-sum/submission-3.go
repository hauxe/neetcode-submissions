func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[target-nums[i]] = i
	}
	for i := range nums {
		if idx, exist := m[nums[i]]; exist && i != idx {
			if i > idx {
				return []int{idx, i}
			}
			return []int{i, idx}
		}
	}
	return nil
}
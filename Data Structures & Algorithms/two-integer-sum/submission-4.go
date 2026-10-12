func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		if idx, exist := m[target-nums[i]]; exist {
			if i > idx {
				return []int{idx, i}
			}
			return []int{i, idx}
		}
		m[nums[i]] = i
	}
	return nil
}
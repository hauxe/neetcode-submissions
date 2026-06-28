
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := range nums {
		m[target-nums[i]] = i
	}
	for i := range nums {
		if idx, exist := m[nums[i]]; exist && idx != i {
			if idx < i {
				return []int{idx, i}
			}
			return []int{i, idx}
		}
	}
	return nil
}
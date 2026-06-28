func hasDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for i := range nums {
		if _, exist := m[nums[i]]; exist {
			return true
		}
		m[nums[i]] = struct{}{}
	}
	return false
}
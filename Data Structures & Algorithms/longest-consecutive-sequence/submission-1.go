func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for i := range nums {
		m[nums[i]] = struct{}{}
	}
	maxLength := 0
	for i := range nums {
		if _, exist := m[nums[i]-1]; exist {
			continue
		}
		l := 1
		for {
			if _, exist := m[nums[i]+l]; !exist {
				break
			}
			l++
		}
		maxLength = max(maxLength, l)
	}
	return maxLength
}
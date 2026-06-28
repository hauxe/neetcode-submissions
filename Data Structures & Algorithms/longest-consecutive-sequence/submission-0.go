func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for i := range nums {
		m[nums[i]] = struct{}{}
	}
	maxLength := 0
	for i := range nums {
		if _, exist := m[nums[i]-1]; exist {
			continue
		}
		len := 1
		for {
			if _, exist := m[nums[i]+len]; !exist {
				break
			}
			len++
		}
		maxLength = max(maxLength, len)
	}
	return maxLength
}
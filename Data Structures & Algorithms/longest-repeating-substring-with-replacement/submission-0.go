func characterReplacement(s string, k int) int {
	m := make(map[byte]int)
	maxLength := 0
	maxCount := 0
	left, right := 0, 0
	for ; right < len(s); right++ {
		m[s[right]]++
		maxCount = max(maxCount, m[s[right]])
		candidate := right - left + 1 - maxCount
		if candidate > k {
			m[s[left]]--
			left++
		}
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}
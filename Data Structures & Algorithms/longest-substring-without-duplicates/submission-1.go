func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxLength := 0
	pivot := -1
	for i := range s {
		if idx, exist := m[s[i]]; exist {
			if idx >= pivot {
				pivot = idx+1
			}
		}
		m[s[i]] = i
		if pivot == -1 {
			pivot = i
		}
		maxLength = max(maxLength, i-pivot+1)
	}
	return maxLength
}

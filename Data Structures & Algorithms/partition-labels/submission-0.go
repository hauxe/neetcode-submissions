func partitionLabels(s string) []int {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]] = max(m[s[i]], i)
	}
	var result []int
	currentIdx := m[s[0]]
	prevIdx := 0
	for i := 1; i <= len(s); i++ {
		if i > currentIdx {
			result = append(result, i-prevIdx)
			prevIdx = i
		}
		if i < len(s) {
			currentIdx = max(i, currentIdx, m[s[i]])
		}
	}
	return result
}
func minWindow(s string, t string) string {
    m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := range t {
		m1[t[i]]++
	}
	minLeft, minRight := -1, -1
	matched := 0
	left, right := 0, 0
	for ; right < len(s); right++ {
		m2[s[right]]++
		if m1[s[right]] > 0 && m1[s[right]] == m2[s[right]] {
			matched++
		}
		for matched == len(m1) {
			if minLeft == -1 || right-left < minRight-minLeft {
				minLeft, minRight = left, right
			}
			m2[s[left]]--
			if m1[s[left]] > 0 && m1[s[left]] > m2[s[left]] {
				matched--
			}
			left++
		}
	}
	if minLeft > -1 {
		return s[minLeft:minRight+1]
	}

	return ""
}

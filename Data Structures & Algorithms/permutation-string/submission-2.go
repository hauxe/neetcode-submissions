func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 0
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := range s1 {
		m1[s1[i]]++
	}
	matched := 0
	for ; right < len(s2); right++ {
		m2[s2[right]]++
		if m1[s2[right]] > 0 && m2[s2[right]] == m1[s2[right]] {
			matched++
		}
		for matched == len(m1) {
			if right-left+1 == len(s1) {
				return true
			}
			m2[s2[left]]--
			if m1[s2[left]] > 0 && m2[s2[left]] < m1[s2[left]] {
				matched--
				left++
				break
			}
			
			left++
		}
		
	}
	return false
}

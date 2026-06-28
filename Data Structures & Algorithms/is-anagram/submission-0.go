func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int, len(s))
	for i := range s {
		m[s[i]]++
	}
	matched := 0
	for i := range t {
		m[t[i]]--
		if m[t[i]] < 0 {
			return false
		} else if m[t[i]] == 0 {
			matched++
		}
	}
	return matched == len(m)
}
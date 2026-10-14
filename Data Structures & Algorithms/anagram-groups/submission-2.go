func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	m := make(map[[26]int]int)
	for i := range strs {
		id := [26]int{}
		for j := range strs[i] {
			id[strs[i][j]-'a']++
		}
		if idx, exist := m[id]; exist {
			result[idx] = append(result[idx], strs[i])
			continue
		}
		result = append(result, []string{strs[i]})
		m[id] = len(result) - 1
	}

	return result
}
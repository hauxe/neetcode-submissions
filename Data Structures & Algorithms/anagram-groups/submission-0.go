func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int]int)
	result := make([][]string, 0)
	for i := range strs {
		id := [26]int{}
		for j := range strs[i] {
			idx := strs[i][j] - 'a'
			id[idx]++
		}
		if idx, exist := m[id]; exist {
			result[idx] = append(result[idx], strs[i])
		} else {
			result = append(result, []string{strs[i]})
			m[id] = len(result) - 1
		}
	}
	return result
}
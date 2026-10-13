func longestCommonPrefix(strs []string) string {
	i := 0
	result := []byte{}

	for {
		if len(strs[0]) <= i {
			return string(result)
		}
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i {
				return string(result)
			}
			if strs[j][i] != strs[0][i] {
				return string(result)
			}
		}
		result = append(result, strs[0][i])
		i++
	}
}
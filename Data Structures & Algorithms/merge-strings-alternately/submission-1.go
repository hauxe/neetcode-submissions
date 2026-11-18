func mergeAlternately(word1 string, word2 string) string {
	n1, n2 := len(word1), len(word2)
	result := make([]byte, 0, n1+n2)
	p1, p2 := 0, 0
	for p1 < n1 && p2 < n2 {
		result = append(result, word1[p1], word2[p2])
		p1++
		p2++
	}
	result = append(result, word1[p1:]...)
	result = append(result, word2[p2:]...)
	return string(result)
}
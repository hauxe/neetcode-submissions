
func countSubstrings(s string) int {
	n := len(s)
	expand := func(left, right int) int {
		count := 0
		for ; left >= 0 && right < n && s[left] == s[right]; left, right = left-1, right+1 {
			count++
		}
		return count
	}
	count := 0
	for i := 0; i < n; i++ {
		count++
		count += expand(i-1, i)
		count += expand(i-1, i+1)
	}
	return count
}
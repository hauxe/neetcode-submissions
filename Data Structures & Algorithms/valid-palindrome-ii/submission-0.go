func validPalindrome(s string) bool {
	begin, end := 0, len(s)-1
	return doCheckValidPalindrome(s, begin, end, 0)

}

func doCheckValidPalindrome(s string, begin, end, count int) bool {
	for begin < end {
		if s[begin] != s[end] {
			if count >= 1 {
				return false
			}
			return doCheckValidPalindrome(s, begin+1, end, count+1) || doCheckValidPalindrome(s, begin, end-1, count+1)
		}
		begin++
		end--
	}
	return true
}
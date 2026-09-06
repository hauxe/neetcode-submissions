func longestPalindrome(s string) string {
	store := make(map[string]string)
	var doGetLongest func(s string) string

	doGetLongest = func(s string) string {
		if len(s) < 2 {
			return s
		}
		if longest, exist := store[s]; exist {
			return longest
		}
		if s[0] == s[len(s)-1] && isPalindromic(s[1:len(s)-1]) {
			store[s] = s
			return s
		}
		s1 := doGetLongest(s[:len(s)-1])
		s2 := doGetLongest(s[1:])
		if len(s2) > len(s1) {
			s1 = s2
		}
		store[s] = s1
		return s1
	}
	return doGetLongest(s)
}

func isPalindromic(s string) bool {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}
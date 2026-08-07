func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	if len(s) == 1 {
		return [][]string{{s}}
	}
	var result [][]string
	for i := 0; i < len(s); i++ {
		if isPalindrome(s[:i+1]) {
			parts := partition(s[i+1:])
			if len(parts) > 0 {
				for j := range parts {
					parts[j] = append(parts[j], s[:i+1])
				}
			} else {
				parts = [][]string{{s[:i+1]}}
			}
			result = append(result, parts...)
		}
	}
	return result
}

func isPalindrome(s string) bool {
	for start, end := 0, len(s)-1; start < end; start, end = start+1, end-1 {
		if s[start] != s[end] {
			return false
		}
	}
	return true
}
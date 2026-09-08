func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	prev2, prev1 := 1, 1 // dp[i-2], dp[i-1]

	for i := 2; i <= len(s); i++ {
		curr := 0
		if s[i-1] != '0' {
			curr += prev1
		}
		tens := int(s[i-2]-'0')*10 + int(s[i-1]-'0')
		if tens >= 10 && tens <= 26 {
			curr += prev2
		}
		prev2, prev1 = prev1, curr
	}

	return prev1
}
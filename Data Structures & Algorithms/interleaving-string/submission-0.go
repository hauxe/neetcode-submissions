func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}
	dp := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			if i+j-1 >= 0 {
				if i-1 >= 0 && s3[i+j-1] == s1[i-1] && dp[i-1][j] {
					dp[i][j] = true
				} else if j-1 >= 0 && s3[i+j-1] == s2[j-1] && dp[i][j-1] {
					dp[i][j] = true
				}
			}
		}
	}

	return dp[n1][n2]
}
var directions = [4][2]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
}

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		best := 1
		for _, d := range directions {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && matrix[ni][nj] > matrix[i][j] {
				best = max(best, 1+dfs(ni, nj))
			}
		}
		memo[i][j] = best
		return best
	}

	maxPath := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxPath = max(maxPath, dfs(i, j))
		}
	}
	return maxPath
}
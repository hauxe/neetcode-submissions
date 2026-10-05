func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	var firstRowZero, firstColZero bool
	for j := range n {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}
	for i := range m {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}
	// mark step
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// process step
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if firstRowZero {
		for j := range n {
			matrix[0][j] = 0
		}
	}
	if firstColZero {
		for i := range m {
			matrix[i][0] = 0
		}
	}
}
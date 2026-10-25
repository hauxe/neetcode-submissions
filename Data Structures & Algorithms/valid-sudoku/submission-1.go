func isValidSudoku(board [][]byte) bool {
	m := make(map[[3]byte]bool)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			rowIdx := [3]byte{'0', byte(i) - '0', board[i][j]}
			colIdx := [3]byte{'1', byte(j) - '0', board[i][j]}
			squareID := (i/3)*3 + j/3
			squareIdx := [3]byte{'2', byte(squareID) - '0', board[i][j]}
			if m[rowIdx] || m[colIdx] || m[squareIdx] {
				fmt.Println()
				fmt.Println()
				return false
			}
			m[rowIdx] = true
			m[colIdx] = true
			m[squareIdx] = true
		}
	}
	return true
}
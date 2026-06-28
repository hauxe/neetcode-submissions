func isValidSudoku(board [][]byte) bool {
	m := make(map[[3]byte]struct{}, 9*3)
	b := byte('.')
	for i := range board {
		for j := range board[i] {
			if board[i][j] == b {
				continue
			}
			rowIdx := [3]byte{'0', byte(j) - '0', board[i][j]}
			if _, exist := m[rowIdx]; exist {
				return false
			}
			m[rowIdx] = struct{}{}
			colIdx := [3]byte{'1', byte(i) - '0', board[i][j]}
			if _, exist := m[colIdx]; exist {
				return false
			}
			m[colIdx] = struct{}{}
			sID := (j/3)*3 + (i / 3)
			squareIdx := [3]byte{'2', byte(sID) - '0', board[i][j]}
			if _, exist := m[squareIdx]; exist {
				return false
			}
			m[squareIdx] = struct{}{}
		}
	}
	return true
}
var directions = [][2]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
	{1, 1}, {1, -1},
	{-1, 1}, {-1, -1},
}

type Solver struct {
	result [][]string
}

func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = 'x'
		}
	}
	solver := &Solver{}
	solver.doSolveNqueens(n, board)
	return solver.result
}

func (s *Solver) doSolveNqueens(n int, board [][]byte) {
	if n == 0 {
		s.result = append(s.result, build(board))
		return
	}
	row := n - 1
	for col := range board[row] {
		if board[row][col] == 'x' {
			board[row][col] = 'Q'
			mark(row, col, board)
			s.doSolveNqueens(n-1, board)
			board[row][col] = 'x'
			sweep(row, col, board)
		}
	}
}

func mark(row int, col int, board [][]byte) {
	for _, direction := range directions {
		for i := range board {
			newRow, newCol := row+direction[0]*(i+1), col+direction[1]*(i+1)
			if newRow >= 0 && newRow < len(board) && newCol >= 0 && newCol < len(board) && board[newRow][newCol] == 'x' {
				board[newRow][newCol] = byte(row) + '0'
			}
		}
	}
}

func sweep(row int, col int, board [][]byte) {
	for _, direction := range directions {
		for i := range board {
			newRow, newCol := row+direction[0]*(i+1), col+direction[1]*(i+1)
			if newRow >= 0 && newRow < len(board) && newCol >= 0 && newCol < len(board) && board[newRow][newCol] == byte(row)+'0' {
				board[newRow][newCol] = 'x'
			}
		}
	}
}

func build(board [][]byte) []string {
	var result []string
	for i := range board {
		b := make([]byte, len(board[i]))
		for j := range board[i] {
			if board[i][j] != 'Q' {
				b[j] = '.'
			} else {
				b[j] = 'Q'
			}
		}
		result = append(result, string(b))
	}
	return result
}
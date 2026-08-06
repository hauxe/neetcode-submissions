var directions = [][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	for row := range board {
		for col := range board[row] {
			if board[row][col] == word[0] {
				if doCheck(row, col, board, word[1:]) {
					return true
				}
			}
		}
	}
	return false
}

func doCheck(row int, col int, board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	for _, direction := range directions {
		newRow, newCol := row+direction[0], col+direction[1]
		if newRow >= 0 && newRow < len(board) && newCol >= 0 && newCol < len(board[0]) && board[newRow][newCol] == word[0] {
			c := board[row][col]
			board[row][col] = '#'
			if doCheck(newRow, newCol, board, word[1:]) {
				return true
			}
			board[row][col] = c
		}
	}
	return false
}
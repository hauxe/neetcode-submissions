var directions = [][2]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
}

func solve(board [][]byte) {
	candidate := make(map[[2]int]int)
	queue := make([][2]int, 0)
	rows, cols := len(board), len(board[0])
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				candidate[[2]int{i, j}] = 1
				if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
					candidate[[2]int{i, j}] = 2
					queue = append(queue, [2]int{i, j})
				}
			}
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, direction := range directions {
			newRow, newCol := cur[0]+direction[0], cur[1]+direction[1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && board[newRow][newCol] == 'O' && candidate[[2]int{newRow, newCol}] == 1 {
				candidate[[2]int{newRow, newCol}] = 2
				queue = append(queue, [2]int{newRow, newCol})
			}
		}
	}
	for item, val := range candidate {
		if val == 1 {
			board[item[0]][item[1]] = 'X'
		}
	}
}
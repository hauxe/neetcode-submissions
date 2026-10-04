func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	directions := [4][6]int{
		// row, col, top, left, bottom, right
		{0, 1, 1, 0, 0, 0},   // right
		{1, 0, 0, 0, 0, -1},  // down
		{0, -1, 0, 0, -1, 0}, // left
		{-1, 0, 0, 1, 0, 0},  // up
	}
	result := make([]int, m*n)
	direction := 0 // left -> right
	row, col := 0, 0
	// boundary
	top, left, bottom, right := 0, 0, m, n
	for i := 0; i < m*n; i++ {
		result[i] = matrix[row][col]
		newRow, newCol := row+directions[direction][0], col+directions[direction][1]
		if newRow >= top && newRow < bottom && newCol >= left && newCol < right {
			row, col = newRow, newCol
		} else {
			top, left, bottom, right = top+directions[direction][2], left+directions[direction][3], bottom+directions[direction][4], right+directions[direction][5]
			direction = (direction + 1) % 4                                       // switch direction
			row, col = row+directions[direction][0], col+directions[direction][1] // always valid based on pre direction order and limitation in outside loop
		}
	}
	return result
}
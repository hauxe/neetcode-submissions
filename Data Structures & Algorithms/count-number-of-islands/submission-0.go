var directions = [][2]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
}

func numIslands(grid [][]byte) int {
	var result int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				mark(i, j, grid)
				
				result++
			}
		}
	}
	return result
}

func mark(row int, col int, grid [][]byte) {
	for _, direction := range directions {
		newRow, newCol := row+direction[0], col+direction[1]
		if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) && grid[newRow][newCol] == '1' {
			grid[newRow][newCol] = '0'
			mark(newRow, newCol, grid)
		}
	}
}
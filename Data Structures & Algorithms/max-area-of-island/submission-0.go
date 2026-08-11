var directions = [][2]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
}

func maxAreaOfIsland(grid [][]int) int {
	var maxArea int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				marked := mark(i, j, grid)
				maxArea = max(maxArea, marked+1)
			}
		}
	}
	return maxArea
}

func mark(row int, col int, grid [][]int) int {
	var marked int
	for _, direction := range directions {
		newRow, newCol := row+direction[0], col+direction[1]
		if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) && grid[newRow][newCol] == 1 {
			grid[newRow][newCol] = 0
			marked++
			marked += mark(newRow, newCol, grid)
		}
	}
	return marked
}
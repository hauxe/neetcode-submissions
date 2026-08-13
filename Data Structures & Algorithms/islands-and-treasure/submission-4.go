var directions = [][2]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
}

func islandsAndTreasure(grid [][]int) {
	queue := [][2]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, direction := range directions {
			newRow, newCol := cur[0]+direction[0], cur[1]+direction[1]
			if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) && grid[newRow][newCol] > grid[cur[0]][cur[1]]+1 {
				grid[newRow][newCol] = grid[cur[0]][cur[1]] + 1
				queue = append(queue, [2]int{newRow, newCol})
			}
		}
	}
}
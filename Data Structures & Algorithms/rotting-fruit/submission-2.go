var directions = [][2]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
}

func orangesRotting(grid [][]int) int {
	fresh := 0
	queue := make([][2]int, 0)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	minutes := 0
	curlen := len(queue)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curlen--
		for _, direction := range directions {
			newRow, newCol := cur[0]+direction[0], cur[1]+direction[1]
			if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) && grid[newRow][newCol] == 1 {
				grid[newRow][newCol] = 2
				fresh--
				queue = append(queue, [2]int{newRow, newCol})
			}
		}
		fmt.Println(curlen)
		if curlen == 0 {
			if len(queue) > 0 {
				minutes++
			}
			curlen = len(queue)
		}
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}
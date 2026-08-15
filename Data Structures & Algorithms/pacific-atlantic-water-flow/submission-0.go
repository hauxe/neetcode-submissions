var directions = [][2]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
}

func pacificAtlantic(heights [][]int) [][]int {
	var result [][]int
	var queue [][4]int
	visited := make(map[[2]int][2]int)

	rows := len(heights)
	cols := len(heights[0])

	for row := range heights {
		for col := range heights[0] {
			if row == 0 || col == 0 || row == rows-1 || col == cols-1 {
				a, p := 0, 0
				if row == 0 || col == 0 {
					p = 1
				}
				if row == rows-1 || col == cols-1 {
					a = 1
				}
				queue = append(queue, [4]int{row, col, a, p})
				if a == 1 && p == 1 {
					result = append(result, []int{row, col})
					visited[[2]int{row, col}] = [2]int{1, 1}
				}
			}
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, direction := range directions {
			newRow, newCol := cur[0]+direction[0], cur[1]+direction[1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && heights[newRow][newCol] >= heights[cur[0]][cur[1]] && (visited[[2]int{newRow, newCol}][0] < cur[2] || visited[[2]int{newRow, newCol}][1] < cur[3]) {
				visited[[2]int{newRow, newCol}] = [2]int{max(visited[[2]int{newRow, newCol}][0], cur[2]), max(visited[[2]int{newRow, newCol}][1], cur[3])}
				if visited[[2]int{newRow, newCol}] == [2]int{1, 1} {
					result = append(result, []int{newRow, newCol})
				}
				queue = append(queue, [4]int{newRow, newCol, visited[[2]int{newRow, newCol}][0], visited[[2]int{newRow, newCol}][1]})
			}
		}
	}

	return result

}
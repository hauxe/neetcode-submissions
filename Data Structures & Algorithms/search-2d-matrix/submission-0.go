func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	begin, end := 0, m*n-1
	for begin <= end {
		middle := (begin+end)/2
		i, j := middle/n, middle%n
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			begin = middle+1
		} else {
			end = middle-1
		}
	}
	return false
}

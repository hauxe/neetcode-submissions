var directions = [][2]int{
    {0, 1}, {0, -1},
    {1, 0}, {-1, 0},
}

func swimInWater(grid [][]int) int {
    n := len(grid)
    if n == 0 {
        return 0
    }

    low := max(grid[0][0], grid[n-1][n-1])
    high := 0
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] > high {
                high = grid[i][j]
            }
        }
    }

    canReach := func(t int) bool {
        if grid[0][0] > t {
            return false
        }
        visited := make([][]bool, n)
        for i := range visited {
            visited[i] = make([]bool, n)
        }
        q := [][2]int{{0, 0}}
        visited[0][0] = true

        for len(q) > 0 {
            r, c := q[0][0], q[0][1]
            q = q[1:]
            if r == n-1 && c == n-1 {
                return true
            }
            for _, d := range directions {
                nr, nc := r+d[0], c+d[1]
                if nr >= 0 && nr < n && nc >= 0 && nc < n &&
                    !visited[nr][nc] && grid[nr][nc] <= t {
                    visited[nr][nc] = true
                    q = append(q, [2]int{nr, nc})
                }
            }
        }
        return false
    }

    for low < high {
        mid := (low + high) / 2
        if canReach(mid) {
            high = mid
        } else {
            low = mid + 1
        }
    }
    return low
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
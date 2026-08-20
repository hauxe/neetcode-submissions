func countComponents(n int, edges [][]int) int {
	graph := make([][]int, n)
	for i := range edges {
		a, b := edges[i][0], edges[i][1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	queue := make([]int, 0)
	visited := make([]bool, n)
	components := 0
	for i := range n {
		if !visited[i] {
			components++
			queue = append(queue, i)
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				visited[cur] = true
				for j := range graph[cur] {
					if visited[graph[cur][j]] {
						continue
					}
					queue = append(queue, graph[cur][j])
				}
			}
		}
	}

	return components

}
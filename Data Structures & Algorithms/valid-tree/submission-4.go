func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	graph := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make([]bool, n)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -1
	}

	queue := []int{0}
	visited[0] = true
	count := 1

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, nei := range graph[cur] {
			if nei == parent[cur] {
				continue
			}
			if visited[nei] {
				return false // cycle found
			}
			visited[nei] = true
			parent[nei] = cur
			count++
			queue = append(queue, nei)
		}
	}

	return count == n // all nodes reached => connected
}
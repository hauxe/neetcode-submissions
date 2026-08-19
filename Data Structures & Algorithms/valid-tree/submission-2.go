func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false // a tree must have exactly n-1 edges
	}

	graph := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a) // undirected: add both directions
	}

	visited := make([]bool, n)
	var hasCycle func(node, parent int) bool
	hasCycle = func(node, parent int) bool {
		visited[node] = true
		for _, nei := range graph[node] {
			if nei == parent {
				continue // skip the edge back to where we came from
			}
			if visited[nei] {
				return true
			}
			if hasCycle(nei, node) {
				return true
			}
		}
		return false
	}

	if hasCycle(0, -1) {
		return false
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			return false // disconnected
		}
	}
	return true
}
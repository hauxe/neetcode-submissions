func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // path compression
		}
		return parent[x]
	}

	for _, e := range edges {
		a, b := e[0], e[1]
		ra, rb := find(a), find(b)
		if ra == rb {
			// a and b already connected -> this edge closes the cycle
			return e
		}
		parent[ra] = rb
	}
	return nil // unreachable given problem constraints
}
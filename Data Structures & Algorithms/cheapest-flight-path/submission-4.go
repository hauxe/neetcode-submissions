func findCheapestPrice(n int, flights [][]int, src, dst, k int) int {
	const inf = math.MaxInt32
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[src] = 0

	for i := 0; i <= k; i++ {
		next := make([]int, n)
		copy(next, dist)
		for _, f := range flights {
			u, v, w := f[0], f[1], f[2]
			if dist[u] != inf && dist[u]+w < next[v] {
				next[v] = dist[u] + w
			}
		}
		dist = next
	}

	if dist[dst] == inf {
		return -1
	}
	return dist[dst]
}
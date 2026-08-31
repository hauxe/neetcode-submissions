func foreignDictionary(words []string) string {
	// Step 1: Initialize graph and in-degree map for all unique letters
	adj := make(map[byte]map[byte]bool)
	inDegree := make(map[byte]int)

	for _, w := range words {
		for i := 0; i < len(w); i++ {
			c := w[i]
			if _, ok := adj[c]; !ok {
				adj[c] = make(map[byte]bool)
				inDegree[c] = 0
			}
		}
	}

	// Step 2: Build graph edges by comparing adjacent words
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := len(w1)
		if len(w2) < minLen {
			minLen = len(w2)
		}

		foundDiff := false
		for j := 0; j < minLen; j++ {
			c1, c2 := w1[j], w2[j]
			if c1 != c2 {
				// Add edge c1 -> c2 if not already present
				if !adj[c1][c2] {
					adj[c1][c2] = true
					inDegree[c2]++
				}
				foundDiff = true
				break
			}
		}

		// Invalid case: w1 is longer than w2 but w2 is a prefix of w1
		if !foundDiff && len(w1) > len(w2) {
			return ""
		}
	}

	// Step 3: Topological sort using Kahn's algorithm (BFS)
	queue := make([]byte, 0)
	for c, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, c)
		}
	}

	result := make([]byte, 0, len(inDegree))

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		result = append(result, c)

		for neighbor := range adj[c] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Step 4: Check if all letters were included (no cycle)
	if len(result) != len(inDegree) {
		return ""
	}

	return string(result)
}
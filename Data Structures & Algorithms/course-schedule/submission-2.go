func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	degree := make([]int, numCourses)
	for _, p := range prerequisites {
		a, b := p[0], p[1]
		graph[b] = append(graph[b], a)
		degree[a]++
	}
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}

	taken := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		taken++
		for i := range graph[cur] {
			course := graph[cur][i]
			degree[course]--
			if degree[course] == 0 {
				queue = append(queue, course)
			}
		}
	}
	return taken == numCourses
}
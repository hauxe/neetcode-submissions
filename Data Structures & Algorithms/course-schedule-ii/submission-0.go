func findOrder(numCourses int, prerequisites [][]int) []int {
	var result []int
	graph := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		graph[a] = append(graph[a], b)
	}
	visiting := 1
	visited := 2

	var runCycle func(course int) bool
	visit := make([]int, numCourses)

	runCycle = func(course int) bool {
		if visit[course] == visited {
			return false
		}
		if visit[course] == visiting {
			return true
		}
		visit[course] = visiting
		for i := range graph[course] {
			if runCycle(graph[course][i]) {
				return true
			}
		}
		visit[course] = visited
		result = append(result, course)
		return false
	}

	for i := range graph {
		if len(graph[i]) == 0 {
			continue
		}
		if runCycle(i) {
			return nil
		}
	}
	for i := range visit {
		if visit[i] == 0 {
			result = append(result, i)
		}
	}
	return result
}
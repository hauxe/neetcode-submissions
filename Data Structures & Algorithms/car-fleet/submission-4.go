func carFleet(target int, position []int, speed []int) int {
	m := make(map[int]int, len(position))
	for i := range position {
		m[position[i]] = speed[i]
	}
	sort.Slice(position, func(i, j int) bool {
		return position[i] > position[j] // descending
	})
	stack := make([]float64, len(position))
	p := 0
	stack[p] = float64(target-position[0]) / float64(m[position[0]])
	for i := 1; i < len(position); i++ {
		prev := stack[p]
		curr := float64(target-position[i]) / float64(m[position[i]])
		if curr > prev {
			p++
			stack[p] = curr
		}
	}
	return p + 1
}
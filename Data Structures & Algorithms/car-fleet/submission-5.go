func carFleet(target int, position []int, speed []int) int {
	m := make(map[int]int, len(position))
	for i := range position {
		m[position[i]] = speed[i]
	}
	sort.Slice(position, func(i, j int) bool {
		return position[i] > position[j] // descending
	})
	p := 0
	prev := float64(target-position[0]) / float64(m[position[0]])
	for i := 1; i < len(position); i++ {
		curr := float64(target-position[i]) / float64(m[position[i]])
		if curr > prev {
			p++
			prev = curr
		}
	}
	return p + 1
}
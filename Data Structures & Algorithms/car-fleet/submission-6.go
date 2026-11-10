func carFleet(target int, position []int, speed []int) int {
	m := make(map[int]int, len(position))
	for i := range position {
		m[position[i]] = speed[i]
	}
	sort.Ints(position)
	stack := make([]int, 0)
	for i := len(position) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			stack = append(stack, position[i])
			continue
		}
		cur := float64(target-stack[len(stack)-1]) / float64(m[stack[len(stack)-1]])
		t := float64(target-position[i]) / float64(m[position[i]])
		if cur < t {
			stack = append(stack, position[i])
		}
	}
	return len(stack)
}
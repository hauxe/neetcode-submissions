func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, len(temperatures))
	p := -1
	for i := range temperatures {
		for ; p >= 0 && temperatures[stack[p]] < temperatures[i]; p-- {
			result[stack[p]] = i - stack[p]
		}
		p++
		stack[p] = i
	}
	return result
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	var last []int
	for _, interval := range intervals {
		if last == nil {
			last = interval
			continue
		}
		if last[1] < interval[0] {
			result = append(result, last)
			last = interval
		} else {
			last = []int{last[0], max(last[1], interval[1])}
		}
	}
	if last != nil {
		result = append(result, last)
	}
	return result
}

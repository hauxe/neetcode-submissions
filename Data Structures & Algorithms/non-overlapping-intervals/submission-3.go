func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	count := 0
	var prev []int
	for _, interval := range intervals {
		if prev == nil {
			prev = interval
			continue
		}
		if prev[0] < interval[1] && prev[1] > interval[0] {
			count++
			if interval[1] < prev[1] {
				fmt.Println("removing", prev)
				prev = interval
			} else {
				fmt.Println("removing", interval)
			}
			continue
		}
		prev = interval

	}
	return count
}
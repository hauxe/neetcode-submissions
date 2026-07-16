func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	startIdx, endIdx := 0, len(intervals)-1
	begin, end := 0, len(intervals)-1
	for begin <= end {
		middle := (begin + end) / 2
		if intervals[middle][0] == newInterval[0] {
			startIdx = middle
			break
		}
		if intervals[middle][0] < newInterval[0] {
			begin = middle + 1
			startIdx = begin
			if intervals[middle][1] >= newInterval[0] {
				startIdx--
				break
			}
		} else {
			end = middle - 1
		}
	}
	fmt.Println("start index", startIdx)
	begin, end = 0, len(intervals)-1
	for begin <= end {
		middle := (begin + end) / 2
		if intervals[middle][1] == newInterval[1] {
			endIdx = middle
			break
		}
		if intervals[middle][1] > newInterval[1] {
			end = middle - 1
			endIdx = end
			if intervals[middle][0] <= newInterval[1] {
				endIdx++
				break
			}
		} else {
			begin = middle + 1
		}
	}
	fmt.Println("end index", endIdx)
	// cut off
	if endIdx == -1 {
		return append([][]int{newInterval}, intervals...)
	} else if startIdx == len(intervals) {
		return append(intervals, newInterval)
	}
	var result [][]int
	result = append(result, intervals[0:startIdx]...)
	result = append(result, []int{min(intervals[startIdx][0], newInterval[0]), max(intervals[endIdx][1], newInterval[1])})
	result = append(result, intervals[endIdx+1:]...)
	return result
}
/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	start, end := make([]int, len(intervals)), make([]int, len(intervals))
	for i := range intervals {
		start[i], end[i] = intervals[i].start, intervals[i].end
	}
	sort.Slice(start, func(i, j int) bool {
		return start[i] < start[j]
	})
	sort.Slice(end, func(i, j int) bool {
		return end[i] < end[j]
	})
	s, e := 0, 0
	count := 0
	minRoom := 0
	for e < len(end) {
		if s < len(start) && start[s] < end[e] {
			s++
			count++
			minRoom = max(minRoom, count)
		} else {
			e++
			count--
		}
	}
	return minRoom
}

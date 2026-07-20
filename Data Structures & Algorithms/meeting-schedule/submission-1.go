/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) == 0 {
		return true
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].end
	})
	lastEnd := intervals[0].end
	for i := 1; i < len(intervals); i++ {
		if lastEnd > intervals[i].start {
			return false
		}
		lastEnd = intervals[i].end
	}
	return true
}

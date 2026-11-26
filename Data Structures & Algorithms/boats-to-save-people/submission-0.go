func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	begin, end := 0, len(people)-1
	boats := 0
	for begin < end {
		sum := people[begin] + people[end]
		switch {
		case people[end] == limit:
			end--
		case sum > limit:
			end--
		case sum <= limit:
			begin++
			end--
		}
		boats++
	}
	if begin == end {
		boats++
	}
	return boats
}
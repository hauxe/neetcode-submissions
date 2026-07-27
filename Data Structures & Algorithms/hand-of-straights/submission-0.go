func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize > 0 {
		return false
	}
	m := make(map[int]int)
	for i := range hand {
		m[hand[i]]++
	}
	sort.Slice(hand, func(i, j int) bool {
		return hand[i] < hand[j]
	})
	for i := range hand {
		if m[hand[i]] == 0 {
			continue
		}
		m[hand[i]]--
		for j := 1; j < groupSize; j++ {
			idx := hand[i] + j
			m[idx]--
			if m[idx] < 0 {
				return false
			}
		}
	}
	return true
}
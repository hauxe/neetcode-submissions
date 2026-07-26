func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	for i := range gas {
		total += (gas[i] - cost[i])
	}
	if total < 0 {
		return -1
	}
	total = 0
	resultIdx := 0
	for i := range gas {
		total += gas[i] - cost[i]
		if total < 0 {
			total = 0
			resultIdx = i + 1
		}
	}
	return resultIdx
}
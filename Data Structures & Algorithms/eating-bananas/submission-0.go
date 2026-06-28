func minEatingSpeed(piles []int, h int) int {
	sum := 0
	maxVal := 0
	for i := range piles {
		sum += piles[i]
		maxVal = max(maxVal, piles[i])
	}
	begin := sum / h
	if sum%h > 0 {
		begin++
	}
	end := maxVal
	minK := 0
	for begin <= end {
		middle := (begin + end) / 2
		candidateHour := 0
		for i := range piles {
			candidateHour += (piles[i] / middle)
			if piles[i]%middle > 0 {
				candidateHour++
			}
		}
		// fmt.Println("begin", begin, "middle", middle, "end", end, candidateHour)
		if candidateHour <= h {
			end = middle - 1
			minK = middle
		} else {
			begin = middle + 1
		}
	}
	return minK
}
func mergeTriplets(triplets [][]int, target []int) bool {
	var matches [3]int
	for i := range triplets {
		if triplets[i][0] > target[0] || triplets[i][1] > target[1] || triplets[i][2] > target[2] {
			continue
		}
		if triplets[i][0] == target[0] {
			matches[0] = 1
		}
		if triplets[i][1] == target[1] {
			matches[1] = 1
		}
		if triplets[i][2] == target[2] {
			matches[2] = 1
		}

		if matches[0] == 1 && matches[1] == 1 && matches[2] == 1 {
			return true
		}
	}
	return false
}
func majorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, x := range nums {
		if count == 0 {
			candidate = x
			count++
			continue
		}
		if x != candidate {
			count--
		} else {
			count++
		}
	}
	return candidate
}
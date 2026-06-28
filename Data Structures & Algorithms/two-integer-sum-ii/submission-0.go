func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		v := numbers[left] + numbers[right]
		if v == target {
			return []int{left+1, right+1}
		}
		if v < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

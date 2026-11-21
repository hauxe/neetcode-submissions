func twoSum(numbers []int, target int) []int {
	begin, end := 0, len(numbers)-1
	for begin < end {
		s := numbers[begin] + numbers[end]
		switch {
		case s == target:
			return []int{begin + 1, end + 1}
		case s > target:
			end--
		default:
			begin++
		}
	}
	return nil
}
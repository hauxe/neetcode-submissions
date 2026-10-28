func majorityElement(nums []int) []int {
	if len(nums) < 3 {
		return nums
	}
	candidate1, candidate2 := 0, 1
	count1, count2 := 0, 0
	for i := range nums {
		switch {
		case candidate1 == nums[i]:
			count1++
		case candidate2 == nums[i]:
			count2++
		case count1 == 0:
			candidate1, count1 = nums[i], 1
		case count2 == 0:
			candidate2, count2 = nums[i], 1
		default:
			count1--
			count2--
		}
	}
	count1, count2 = 0, 0
	for i := range nums {
		switch nums[i] {
		case candidate1:
			count1++
		case candidate2:
			count2++
		}
	}
	result := []int{}
	n := len(nums) / 3
	if count1 > n {
		result = append(result, candidate1)
	}
	if count2 > n {
		result = append(result, candidate2)
	}
	return result
}
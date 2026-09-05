func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob2(nums[1:]), rob2(nums[:len(nums)-1]))
}

func rob2(nums []int) int {
	var doRob func(i int) int
	robbed := make(map[int]int)
	doRob = func(i int) int {
		if i >= len(nums) {
			return 0
		}
		if i == len(nums)-1 {
			return nums[i]
		}
		if val, exist := robbed[i]; exist {
			return val
		}
		rob1 := doRob(i + 1)
		rob2 := doRob(i+2) + nums[i]
		robbed[i] = max(rob1, rob2)

		return robbed[i]
	}
	return doRob(0)
}
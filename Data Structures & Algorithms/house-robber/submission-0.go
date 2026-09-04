func rob(nums []int) int {
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
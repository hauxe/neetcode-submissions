func maxSlidingWindow(nums []int, k int) []int {
    queue := make([]int, len(nums))
	result := make([]int, 0, len(nums)-k)
	head := 0
	tail := -1
	for i := range nums {
		if tail >= 0 && queue[head] < i+1-k {
			
			head++
		}
		for ; tail >= head; tail-- {
			if nums[queue[tail]] >= nums[i] {
				break
			}
		}
		tail++
		queue[tail] = i
		if i+1 >= k {
			result = append(result, nums[queue[head]])
		}
	}
	return result
}

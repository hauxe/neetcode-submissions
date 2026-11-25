func maxArea(height []int) int {
	var maxValue int
	left, right := 0, len(height)-1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		maxValue = max(maxValue, area)
		if height[left] >= height[right] {
			right--
		} else {
			left++
		}
	}
	return maxValue
}
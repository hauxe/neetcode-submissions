func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]
	trapped := 0
	for left <= right {
		if maxLeft <= maxRight {
			trapped += max(0, maxLeft-height[left])
			maxLeft = max(maxLeft, height[left])
			left++
		} else {
			trapped += max(0, maxRight-height[right])
			maxRight = max(maxRight, height[right])
			right--
		}
	}
	return trapped
}
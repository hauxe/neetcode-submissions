func trap(height []int) int {
	left, right := 0, len(height) - 1
	bestLeft, bestRight := -1, -1
	trapped := 0
	for left < right {
		if height[left] < height[right] {
			// move left
			if bestLeft == -1 || bestLeft <= height[left] {
				bestLeft = height[left]
			} else {
				trapped += (bestLeft - height[left])
			}
			left++
		} else {
			// move right
			if bestRight == -1 || bestRight <= height[right] {
				bestRight = height[right]
			} else {
				trapped += (bestRight - height[right])
			}
			right--
		}
	}
	return trapped
}

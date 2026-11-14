func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, 0)
	var maxArea int
	// left -> right: build increasing stack
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			width := i - left - 1
			maxArea = max(maxArea, heights[top]*width)
		}
		stack = append(stack, i)
	}
	// right -> left: build max area from end
	end := n
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left := -1
		if len(stack) > 0 {
			left = stack[len(stack)-1]
		}
		width := end - left - 1
		maxArea = max(maxArea, heights[top]*width)
	}
	return maxArea
}
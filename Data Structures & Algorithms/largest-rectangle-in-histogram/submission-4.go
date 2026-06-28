func largestRectangleArea(heights []int) int {
	type entry struct{ idx, left int }
	stack := make([]entry, len(heights))
	p := -1
	maxArea := 0
	for i, h := range heights {
		left := i
		for ; p >= 0 && heights[stack[p].idx] > h; p-- {
			top := stack[p]
			area := heights[top.idx] * (i - top.left)
			maxArea = max(area, maxArea)
			left = top.left
		}
		p++
		stack[p] = entry{idx: i, left: left}
	}
	end := len(heights)
	for j := p; j >= 0; j-- {
		top := stack[j]
		area := heights[top.idx] * (end - top.left)
		maxArea = max(area, maxArea)
	}
	return maxArea
}
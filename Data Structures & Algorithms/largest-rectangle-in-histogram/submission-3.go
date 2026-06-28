func largestRectangleArea(heights []int) int {
    type entry struct{ idx, left int }
    stack := make([]entry, 0, len(heights))
    maxArea := 0

    for i, h := range heights {
        left := i
        for len(stack) > 0 && heights[stack[len(stack)-1].idx] > h {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            area := heights[top.idx] * (i - top.left)
            if area > maxArea {
                maxArea = area
            }
            left = top.left // inherit left boundary from the popped bar
        }
        stack = append(stack, entry{i, left})
    }

    // flush remaining bars — all extend to the end
    n := len(heights)
    for _, e := range stack {
        area := heights[e.idx] * (n - e.left)
        if area > maxArea {
            maxArea = area
        }
    }
    return maxArea
}
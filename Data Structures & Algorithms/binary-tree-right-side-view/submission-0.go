/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	levels := levelOrder(root)
	rightSideView := []int{}
	for i := range levels {
		l := len(levels[i]) - 1
		rightSideView = append(rightSideView, levels[i][l])
	}
	return rightSideView
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{[]int{root.Val}}
	left, right := levelOrder(root.Left), levelOrder(root.Right)
	leftLen, rightLen := len(left), len(right)
	l := max(leftLen, rightLen)
	for i := 0; i < l; i++ {
		combined := []int{}
		if i < leftLen {
			combined = append(combined, left[i]...)
		}
		if i < rightLen {
			combined = append(combined, right[i]...)
		}
		result = append(result, combined)
	}
	return result
}

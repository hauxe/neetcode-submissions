/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	left, right := rightSideView(root.Left), rightSideView(root.Right)
	leftLen, rightLen := len(left), len(right)
	l := max(leftLen, rightLen)
	for i := 0; i < l; i++ {
		if i < rightLen {
			result = append(result, right[i])
			continue
		}
		if i < leftLen {
			result = append(result, left[i])
		}
	}
	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(float64(longest(root.Left)-longest(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func longest(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(longest(root.Left), longest(root.Right)) + 1
}

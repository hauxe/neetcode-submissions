/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(longest(root.Left)+longest(root.Right), diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}

func longest(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(longest(root.Left), longest(root.Right)) + 1
}

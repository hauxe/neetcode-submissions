/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	val, _ := checkKthSmallest(root, k)
	return val
}

func checkKthSmallest(root *TreeNode, k int) (int, int) {
	if root == nil {
		return 0, k
	}
	val, left := checkKthSmallest(root.Left, k)
	if left == -1 {
		return val, -1
	}
	if left == 1 {
		return root.Val, -1
	}
	return checkKthSmallest(root.Right, left-1)
}

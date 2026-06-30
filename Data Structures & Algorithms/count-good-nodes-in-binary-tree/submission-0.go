/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	return countGoodNodes(root, nil)
}

func countGoodNodes(root *TreeNode, maxNode *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	if maxNode == nil || root.Val >= maxNode.Val {
		count = 1
	}
	if maxNode == nil || root.Val > maxNode.Val {
		maxNode = root
	}
	return count + countGoodNodes(root.Left, maxNode) + countGoodNodes(root.Right, maxNode)
}

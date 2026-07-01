/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return checkValidBST(root, nil, nil)
}

func checkValidBST(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil {
		if min != nil && root.Left.Val <= min.Val {
			return false
		}
		if root.Left.Val >= root.Val {
			return false
		}
	}
	if root.Right != nil {
		if max != nil && root.Right.Val >= max.Val {
			return false
		}
		if root.Right.Val <= root.Val {
			return false
		}
	}
	return checkValidBST(root.Left, min, root) && checkValidBST(root.Right, root, max)
}

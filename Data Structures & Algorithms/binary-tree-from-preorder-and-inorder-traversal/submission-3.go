/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	var inorderLeft, inorderRight []int
	for i := range inorder {
		if inorder[i] == rootVal {
			inorderLeft = inorder[0:i]
			inorderRight = inorder[i+1:]
			break
		}
	}
	preorderLeft := preorder[1 : 1+len(inorderLeft)]
	preorderRight := preorder[1+len(inorderLeft) : 1+len(inorderLeft)+len(inorderRight)]

	
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorderLeft, inorderLeft),
		Right: buildTree(preorderRight, inorderRight),
	}
}

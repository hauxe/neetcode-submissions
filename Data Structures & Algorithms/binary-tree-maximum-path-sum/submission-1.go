/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	maxRoot, maxRootCandidate := checkMaxPathSum(root)
	return max(maxRoot, maxRootCandidate)
}


func checkMaxPathSum(root *TreeNode) (int, int) {
	maxLeft, candidateLeft, maxRight, candidateRight := math.MinInt, math.MinInt, math.MinInt, math.MinInt
	sumLeft := 0
	sumRight := 0
	if root.Left != nil {
		maxLeft, candidateLeft = checkMaxPathSum(root.Left)
		sumLeft = maxLeft
	}
	if root.Right != nil {
		maxRight, candidateRight = checkMaxPathSum(root.Right)
		sumRight = maxRight
	}
	maxRoot := max(root.Val, sumLeft+root.Val, sumRight+root.Val)

	maxRootCandidate := max(maxLeft, maxRight, candidateLeft, candidateRight, sumLeft+sumRight+root.Val)
	return maxRoot, maxRootCandidate
}

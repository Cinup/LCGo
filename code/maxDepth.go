package code

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		return max(leftDepth, rightDepth) + 1
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

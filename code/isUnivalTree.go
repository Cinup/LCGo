package code

func isUnivalTree(root *TreeNode) bool {
	return compare(root, root.Val)
}
func compare(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val == val {
		return compare(root.Left, val) && compare(root.Right, val)
	} else {
		return false
	}
}

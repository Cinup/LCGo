package code

func isSymmetric(root *TreeNode) bool {
	return compare(root, root)
}
func compare(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val &&
		compare(t1.Left, t2.Right) &&
		compare(t1.Right, t2.Left)
}

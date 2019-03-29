package code

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	//叶子结点
	if root.Left == nil && root.Right == nil {
		return (sum - root.Val) == 0
	}
	return hasPathSum(root.Left, sum-root.Val) ||
		hasPathSum(root.Right, sum-root.Val)
}

package code

/*
* @Author:cp
* @Parameter:树的根结点,目标
* @Function:树的路径上是否存在和为sum的路径
* @Return:判断结果
*/
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

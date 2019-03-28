package code

/*
* @Author:cp
* @Parameter:二叉树的根结点
* @Function:交换二叉树的左右结点
* @Return:二叉树的根结点
*/
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}

package code

/*
* @Author:cp
* @Parameter:二叉树的根结点
* @Function:二叉树的中序遍历
* @Return:遍历数组
*/
//全局变量存储遍历结果
var result []int

func inorderTraversal(root *TreeNode) []int {
	//每次调用前需要清空,不然会保存上次遍历的结果
	result = nil
	result = traversal(root)
	return result
}
func traversal(root *TreeNode) []int {
	if root != nil {
		traversal(root.Left)
		result = append(result, root.Val)
		traversal(root.Right)
	}
	return result
}

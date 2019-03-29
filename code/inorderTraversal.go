package code

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

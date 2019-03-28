package code

/*
* @Author:cp
* @Parameter:两个树的根结点
* @Function:合并两个树,相同位置的数字相加
* @Return:合并后的树
*/
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

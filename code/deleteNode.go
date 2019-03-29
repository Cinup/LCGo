package code

func deleteNode(node *ListNode) {
	//不删除尾节点
	if node.Next == nil {
		return
	}
	//直接用下一个节点覆盖当前节点
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

package code

/*
* @Author:cp
* @Parameter:链表头节点,需要移除的元素
* @Function:移除链表中指定的元素
* @Return:移除元素后的链表头节点
*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	//设立虚拟头节点,解决删除的是头节点的问题
	tmpHead := new(ListNode)
	tmpHead.Next = head
	current := tmpHead
	for head != nil {
		if head.Val != val {
			current.Next = head
			current = current.Next
		}
		head = head.Next

	}
	current.Next = nil
	return tmpHead.Next
}

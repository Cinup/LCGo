package code

/*
* @Author:cp
* @Parameter:链表头节点
* @Function:找到链表的中间节点
* @Return:中间节点
*/
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for {
		if fast.Next == nil {
			return slow
		}
		if fast.Next.Next == nil {
			return slow.Next
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
}

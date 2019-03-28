package code

/*
* @Author:cp
* @Parameter:链表的头节点
* @Function:判断链表中是否有环
* @Return:判断结果
*/
func hasCycle(head *ListNode) bool {
	//1个节点以下构不成环
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		//一个节点跑一步,一个节点跑两步,如果存在环就会相遇,反之会到达链表末尾
		slow = slow.Next
		fast = fast.Next.Next
		if fast == nil || fast.Next == nil {
			return false
		}

	}
	return true
}

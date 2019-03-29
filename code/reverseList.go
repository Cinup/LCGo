package code

func reverseList(head *ListNode) *ListNode {
	//小于1个节点无需反转
	if head == nil {
		return head
	}
	pre := head
	next := head.Next
	pre.Next = nil
	for next != nil {
		tmpNext := next.Next
		next.Next = pre
		pre = next
		next = tmpNext
	}
	return pre
}

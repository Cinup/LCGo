package code

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	current := head.Val
	cursor := head
	for tmp := head.Next; tmp != nil; {
		if tmp.Val == current {
			tmp = tmp.Next
		} else {
			current = tmp.Val
			cursor.Next = tmp
			cursor = tmp
			tmp = tmp.Next
		}
	}
	cursor.Next = nil
	return head
}

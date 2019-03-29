package code

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

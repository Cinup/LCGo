package code

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	list := new(ListNode)
	tmp := list
	for {
		if l1.Val < l2.Val {
			tmp.Val = l1.Val
			tmp.Next = new(ListNode)
			tmp = tmp.Next
			l1 = l1.Next
			if l1 == nil {
				tmp.Val = l2.Val
				tmp.Next = l2.Next
				return list
			}
		} else {
			tmp.Val = l2.Val
			tmp.Next = new(ListNode)
			tmp = tmp.Next
			l2 = l2.Next
			if l2 == nil {
				tmp.Val = l1.Val
				tmp.Next = l1.Next
				return list
			}
		}
	}
}

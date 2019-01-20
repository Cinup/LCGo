package code

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	list := new(ListNode)
	tmp := list
	for{
		if l1.Val < l2.Val {
			tmp.Val = l1.Val
			tmp.Next = new(ListNode)
			tmp = tmp.Next
			l1 = l1.Next
			if l1 == nil {
				list.Next = l2
				return list
			}
		}else {
			list.Val = l2.Val
			tmp.Next = new(ListNode)
			tmp = tmp.Next
			l2 = l2.Next
			if l2 == nil {
				list.Next = l1
				return list
			}
		}
	}
	//return list
}
